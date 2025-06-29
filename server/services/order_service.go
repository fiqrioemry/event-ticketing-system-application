package services

import (
	"os"
	"server/dto"
	customErr "server/errors"
	"server/models"
	"server/repositories"
	"server/utils"
	"time"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"gorm.io/gorm"
)

type OrderService interface {
	GetOrderDetail(orderID string) ([]dto.OrderDetailResponse, error)
	GetUserTicketsByOrder(orderID string, userID string) ([]dto.UserTicketResponse, error)
	RefundOrder(orderID string, userID string, reason string) (*dto.RefundOrderResponse, error)
	CreateNewOrder(req dto.CreateOrderRequest, userID string) (*dto.CheckoutSessionResponse, error)
	GetMyOrders(userID string, params dto.OrderQueryParams) ([]dto.OrderResponse, *dto.PaginationResponse, error)
}

type orderService struct {
	repo       repositories.OrderRepository
	user       repositories.UserRepository
	ticket     repositories.TicketRepository
	event      repositories.EventRepository
	userTicket repositories.UserTicketRepository
}

func NewOrderService(repo repositories.OrderRepository, user repositories.UserRepository, ticket repositories.TicketRepository, event repositories.EventRepository, userTicket repositories.UserTicketRepository) OrderService {
	return &orderService{repo, user, ticket, event, userTicket}
}

func (s *orderService) CreateNewOrder(req dto.CreateOrderRequest, userID string) (*dto.CheckoutSessionResponse, error) {
	var result *dto.CheckoutSessionResponse

	_, err := s.repo.WithTx(func(tx *gorm.DB) (string, error) {
		user, err := s.user.GetUserByID(userID)
		if err != nil {
			return "", customErr.NewNotFound("user not found")
		}

		event, err := s.event.GetEventByID(req.EventID)
		if err != nil {
			return "", customErr.NewInternal("failed to get event", err)
		}

		if event == nil {
			return "", customErr.NewNotFound("event not found")
		}
		if event.Status != "active" && event.Status != "ongoing" {
			return "", customErr.NewBadRequest("event is not available for ordering")
		}

		orderID := uuid.New()
		order := &models.Order{
			ID:         orderID,
			UserID:     user.ID,
			EventID:    event.ID,
			Fullname:   req.Fullname,
			Email:      req.Email,
			Phone:      req.Phone,
			TotalPrice: 0,
			Status:     "pending",
		}

		var totalPrice float64
		var stripeItems []*stripe.CheckoutSessionLineItemParams

		for _, item := range req.OrderDetails {
			ticket, err := s.ticket.GetTicketByID(item.TicketID)
			if err != nil {
				return "", customErr.NewNotFound("ticket not found")
			}
			if ticket.Quota < item.Quantity {
				return "", customErr.NewBadRequest("not enough quota for ticket: " + ticket.Name)
			}
			if ticket.Limit < item.Quantity {
				return "", customErr.NewBadRequest("ticket limit exceeded for: " + ticket.Name)
			}

			ticket.Quota -= item.Quantity
			if err := tx.Save(ticket).Error; err != nil {
				return "", customErr.NewInternal("failed to update ticket quota", err)
			}

			subtotal := ticket.Price * float64(item.Quantity)
			totalPrice += subtotal

			orderDetail := &models.OrderDetail{
				ID:         uuid.New(),
				OrderID:    orderID,
				TicketID:   ticket.ID,
				TicketName: ticket.Name,
				Quantity:   item.Quantity,
				Price:      ticket.Price,
			}
			if err := tx.Create(orderDetail).Error; err != nil {
				return "", customErr.NewInternal("failed to create order detail", err)
			}

			stripeItems = append(stripeItems, &stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("idr"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(ticket.Name),
					},
					UnitAmount: stripe.Int64(int64(ticket.Price * 100)),
				},
				Quantity: stripe.Int64(int64(item.Quantity)),
			})
		}

		order.TotalPrice = totalPrice
		if err := tx.Create(order).Error; err != nil {
			return "", customErr.NewInternal("failed to create order", err)
		}

		paymentID := uuid.New()
		payment := &models.Payment{
			ID:       paymentID,
			UserID:   user.ID,
			OrderID:  order.ID,
			Fullname: req.Fullname,
			Email:    req.Email,
			Method:   "stripe",
			Status:   "pending",
			Amount:   totalPrice,
		}
		if err := tx.Create(payment).Error; err != nil {
			return "", customErr.NewInternal("failed to create payment", err)
		}

		params := &stripe.CheckoutSessionParams{
			PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
			LineItems:          stripeItems,
			Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
			SuccessURL:         stripe.String(os.Getenv("STRIPE_SUCCESS_URL_DEV")),
			CancelURL:          stripe.String(os.Getenv("STRIPE_CANCEL_URL_DEV")),
			ClientReferenceID:  stripe.String(order.ID.String()),
			Metadata: map[string]string{
				"user_id":    user.ID.String(),
				"order_id":   order.ID.String(),
				"payment_id": paymentID.String(),
			},
		}
		sess, err := session.New(params)
		if err != nil {
			return "", customErr.NewInternal("failed to create stripe session", err)
		}

		result = &dto.CheckoutSessionResponse{
			PaymentID: paymentID.String(),
			SessionID: sess.ID,
			URL:       sess.URL,
		}
		return order.ID.String(), nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *orderService) GetMyOrders(userID string, params dto.OrderQueryParams) ([]dto.OrderResponse, *dto.PaginationResponse, error) {
	orders, total, err := s.repo.GetMyOrders(userID, params)
	if err != nil {
		return nil, nil, err
	}

	var results []dto.OrderResponse
	for _, o := range orders {
		results = append(results, dto.OrderResponse{
			ID:         o.ID.String(),
			EventName:  o.Event.Title,
			EventImage: o.Event.Image,
			EventID:    o.Event.ID.String(),
			Fullname:   o.Fullname,
			Email:      o.Email,
			Phone:      o.Phone,
			TotalPrice: o.TotalPrice,
			Status:     o.Status,
			CreatedAt:  o.CreatedAt.Format("2006-01-02"),
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return results, pagination, nil
}

func (s *orderService) GetOrderDetail(orderID string) ([]dto.OrderDetailResponse, error) {
	orderDetails, err := s.repo.GetOrderDetails(orderID)
	if err != nil {
		return nil, customErr.NewInternal("failed to get order details", err)
	}
	if len(orderDetails) == 0 {
		return nil, customErr.NewNotFound("order detail not found")
	}

	var responses []dto.OrderDetailResponse
	for _, detail := range orderDetails {
		responses = append(responses, dto.OrderDetailResponse{
			ID:         detail.ID.String(),
			TicketName: detail.TicketName,
			TicketID:   detail.TicketID.String(),
			Quantity:   detail.Quantity,
			Price:      detail.Price,
			CreatedAt:  detail.CreatedAt.Format("2006-01-02"),
		})
	}

	return responses, nil
}

func (s *orderService) GetUserTicketsByOrder(orderID, userID string) ([]dto.UserTicketResponse, error) {

	order, err := s.repo.GetOrderByID(orderID)
	if order == nil {
		return nil, customErr.NewNotFound("order not found")
	}

	if err != nil {
		return nil, customErr.NewInternal("failed to get order", err)
	}

	userTickets, err := s.userTicket.GetUserTickets(order.EventID.String(), userID)
	if err != nil {
		return nil, customErr.NewInternal("failed to get user tickets", err)
	}

	var result []dto.UserTicketResponse
	for _, ticket := range userTickets {
		result = append(result, dto.UserTicketResponse{
			ID:         ticket.ID.String(),
			TicketID:   ticket.TicketID.String(),
			EventID:    ticket.EventID.String(),
			EventName:  ticket.Event.Title,
			TicketName: ticket.Ticket.Name,
			QRCode:     ticket.QRCode,
			IsUsed:     ticket.IsUsed,
		})
	}

	return result, nil
}

func (s *orderService) RefundOrder(orderID string, userID string, reason string) (*dto.RefundOrderResponse, error) {
	order, err := s.repo.GetOrderByID(orderID)
	if err != nil || order == nil {
		return nil, customErr.NewNotFound("order not found")
	}
	if order.UserID.String() != userID {
		return nil, customErr.NewForbidden("not your order")
	}
	if order.Status != "paid" || order.IsRefunded {
		return nil, customErr.NewBadRequest("order not refundable")
	}

	eventDate := order.Event.Date
	// prevent refund on event day
	dayNow := time.Now().In(time.Local).Format("2006-01-02")
	dayEvent := eventDate.Format("2006-01-02")
	if dayNow == dayEvent {
		return nil, customErr.NewBadRequest("cannot refund on event day")
	}

	details, _ := s.repo.GetOrderDetails(orderID)
	totalRefund := 0.0

	for _, d := range details {
		ticket, err := s.ticket.GetTicketByID(d.TicketID.String())

		if err != nil {
			return nil, customErr.NewInternal("failed to get ticket", err)
		}

		if ticket == nil {
			return nil, customErr.NewNotFound("ticket not found: ")
		}

		if !ticket.Refundable {
			return nil, customErr.NewBadRequest("ticket not refundable: " + ticket.Name)
		}

		percent := float64(ticket.RefundPercent) / 100.0
		refAmt := float64(d.Quantity) * d.Price * percent
		totalRefund += refAmt
	}

	used, err := s.repo.HasUsedTicket(orderID)
	if err != nil {
		return nil, customErr.NewInternal("failed to check user tickets", err)
	}
	if used {
		return nil, customErr.NewBadRequest("you have already used one of the tickets")
	}

	// mark order
	now := time.Now()
	order.Status = "refunded"
	order.IsRefunded = true
	order.RefundedAt = &now
	order.RefundReason = reason
	order.RefundAmount = totalRefund
	if err := s.repo.UpdateOrder(order); err != nil {
		return nil, err
	}

	// update payment
	s.repo.UpdatePaymentStatus(order.ID.String(), "refunded")
	// update user balance
	s.repo.IncreaseUserBalance(userID, totalRefund)

	user, _ := s.user.GetUserByID(userID)
	return &dto.RefundOrderResponse{
		OrderID:      order.ID.String(),
		RefundAmount: totalRefund,
		RefundedAt:   now.Format(time.RFC3339),
		UserBalance:  user.Balance,
	}, nil
}
