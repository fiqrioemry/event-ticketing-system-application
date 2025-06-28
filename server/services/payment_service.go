package services

import (
	"encoding/json"
	"fmt"
	"server/models"
	"server/repositories"
	"time"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v75"
)

type PaymentService interface {
	ExpireOldPendingPayments() error
	StripeWebhookNotification(event stripe.Event) error
}
type paymentService struct {
	repo       repositories.PaymentRepository
	order      repositories.OrderRepository
	ticket     repositories.TicketRepository
	userTicket repositories.UserTicketRepository
}

func NewPaymentService(repo repositories.PaymentRepository, order repositories.OrderRepository, ticket repositories.TicketRepository, userTicket repositories.UserTicketRepository) PaymentService {
	return &paymentService{repo, order, ticket, userTicket}
}

// ** khusus cron job update status to failed
func (s *paymentService) ExpireOldPendingPayments() error {

	rows, err := s.repo.ExpireOldPendingPayments()
	if err != nil {
		return fmt.Errorf("failed to expire pending payments: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("no expired pending payments found")
	}
	fmt.Printf("%d pending payments marked as failed\n", rows)
	return nil
}

func (s *paymentService) StripeWebhookNotification(event stripe.Event) error {
	if event.Type != "checkout.session.completed" {
		return fmt.Errorf("%s is not a valid event", event.Type)
	}

	var session stripe.CheckoutSession
	if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
		return fmt.Errorf("invalid session data")
	}

	paymentID, ok := session.Metadata["payment_id"]

	if !ok || paymentID == "" {
		return fmt.Errorf("missing order_id in Stripe metadata")
	}

	payment, err := s.repo.GetPaymentByID(paymentID)
	if err != nil {
		return fmt.Errorf("payment query failed: %w", err)
	}
	if payment == nil {
		return fmt.Errorf("payment not found")
	}

	if payment.Status == "paid" {
		return nil
	}
	payment.Method = "card"
	payment.Status = "paid"
	now := time.Now().UTC()
	err = s.repo.UpdatePayment(&models.Payment{
		ID:     payment.ID,
		Method: "card",
		Status: "paid",
		PaidAt: &now,
	})
	if err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}
	order, err := s.order.GetOrderByID(payment.OrderID.String())
	if err != nil || order == nil {
		return fmt.Errorf("order not found: %w", err)
	}

	if err := s.order.UpdateOrderStatus(order.ID.String(), "paid"); err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	orderDetails, err := s.order.GetOrderDetails(order.ID.String())
	if err != nil {
		return fmt.Errorf("failed to fetch order details: %w", err)
	}

	for _, detail := range orderDetails {
		ticket, err := s.ticket.GetTicketByID(detail.TicketID.String())
		if err != nil || ticket == nil {
			return fmt.Errorf("failed to fetch ticket: %w", err)
		}

		ticket.Sold += detail.Quantity

		if err := s.ticket.UpdateTicket(ticket); err != nil {
			return fmt.Errorf("failed to update ticket sold count: %w", err)
		}
	}

	for _, detail := range orderDetails {
		existingCount, err := s.userTicket.CountUserTicketsByTicketID(detail.TicketID)
		if err != nil {
			return fmt.Errorf("failed to count existing user tickets: %w", err)
		}

		for i := 0; i < detail.Quantity; i++ {
			qrCode := fmt.Sprintf("TICKET-%s-%d", detail.TicketID.String(), existingCount+int64(i)+1)

			userTicket := &models.UserTicket{
				ID:       uuid.New(),
				UserID:   order.UserID,
				EventID:  order.EventID,
				TicketID: detail.TicketID,
				IsUsed:   false,
				QRCode:   qrCode,
			}

			if err := s.userTicket.CreateUserTicket(userTicket); err != nil {
				return fmt.Errorf("failed to create user ticket: %w", err)
			}
		}
	}

	return nil

}
