package services

import (
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/config"
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"

	"github.com/fiqrioemry/go-api-toolkit/response"
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
	GetMyOrders(userID string, params dto.OrderQueryParams) ([]dto.OrderResponse, int, error)
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
		if user == nil || err != nil {
			return "", response.NewNotFound("user not found")
		}

		event, err := s.event.GetEventByID(req.EventID)
		if event == nil || err != nil {
			return "", response.NewNotFound("event not found")
		}

		if event.Status != "active" && event.Status != "ongoing" {
			return "", response.NewBadRequest("event is not available for ordering")
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
			if ticket == nil || err != nil {
				return "", response.NewNotFound("ticket not found: " + item.TicketID)
			}

			if ticket.Quota < item.Quantity {
				return "", response.NewBadRequest("not enough quota for ticket: " + ticket.Name)
			}
			if ticket.Limit < item.Quantity {
				return "", response.NewBadRequest("ticket limit exceeded for: " + ticket.Name)
			}

			ticket.Quota -= item.Quantity
			if err := tx.Save(ticket).Error; err != nil {
				return "", response.NewInternalServerError("failed to update ticket quota", err)
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
				return "", response.NewInternalServerError("failed to create order detail", err)
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
			return "", response.NewInternalServerError("failed to create order", err)
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
			return "", response.NewInternalServerError("failed to create payment", err)
		}

		params := &stripe.CheckoutSessionParams{
			PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
			LineItems:          stripeItems,
			Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
			SuccessURL:         stripe.String(config.AppConfig.StripeSuccessUrlDev),
			CancelURL:          stripe.String(config.AppConfig.StripeCancelUrlDev),
			ClientReferenceID:  stripe.String(order.ID.String()),
			Metadata: map[string]string{
				"user_id":    user.ID.String(),
				"order_id":   order.ID.String(),
				"payment_id": paymentID.String(),
			},
		}
		sess, err := session.New(params)
		if err != nil {
			return "", response.NewInternalServerError("failed to create stripe session", err)
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

func (s *orderService) GetMyOrders(userID string, params dto.OrderQueryParams) ([]dto.OrderResponse, int, error) {
	orders, total, err := s.repo.GetMyOrders(userID, params)
	if err != nil {
		return nil, 0, response.NewInternalServerError("failed to retrieve orders", err)
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
			CreatedAt:  o.CreatedAt,
		})
	}

	return results, int(total), nil
}

func (s *orderService) GetOrderDetail(orderID string) ([]dto.OrderDetailResponse, error) {
	orderDetails, err := s.repo.GetOrderDetails(orderID)
	if err != nil || len(orderDetails) == 0 {
		return nil, response.NewNotFound("order details not found").WithContext("orderID", orderID)
	}

	var responses []dto.OrderDetailResponse
	for _, detail := range orderDetails {
		responses = append(responses, dto.OrderDetailResponse{
			ID:         detail.ID.String(),
			TicketName: detail.TicketName,
			TicketID:   detail.TicketID.String(),
			Quantity:   detail.Quantity,
			Price:      detail.Price,
			CreatedAt:  detail.CreatedAt,
		})
	}

	return responses, nil
}

func (s *orderService) GetUserTicketsByOrder(orderID, userID string) ([]dto.UserTicketResponse, error) {

	order, err := s.repo.GetOrderByID(orderID)
	if order == nil || err != nil {
		return nil, response.NewNotFound("order not found").WithContext("orderID", orderID)
	}

	userTickets, err := s.userTicket.GetUserTickets(order.EventID.String(), userID)
	if err != nil || len(userTickets) == 0 {
		return nil, response.NewNotFound("user tickets not found").WithContext("orderID", orderID)
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
		return nil, response.NewNotFound("order not found").WithContext("orderID", orderID)
	}

	if order.UserID.String() != userID {
		return nil, response.NewForbidden("not your order")
	}

	if order.Status != "paid" || order.IsRefunded {
		return nil, response.NewBadRequest("order not refundable")
	}

	eventDate := order.Event.Date
	dayNow := time.Now().In(time.Local)
	dayEvent := eventDate
	if dayNow.Equal(dayEvent) {
		return nil, response.NewBadRequest("cannot refund on event day")
	}

	details, _ := s.repo.GetOrderDetails(orderID)
	totalRefund := 0.0

	for _, d := range details {
		ticket, err := s.ticket.GetTicketByID(d.TicketID.String())
		if ticket == nil || err != nil {
			return nil, response.NewNotFound("ticket not found: "+d.TicketID.String()).WithContext("orderID", orderID)
		}

		if !ticket.Refundable {
			return nil, response.NewBadRequest("ticket not refundable: " + ticket.Name)
		}

		percent := float64(ticket.RefundPercent) / 100.0
		refAmt := float64(d.Quantity) * d.Price * percent
		totalRefund += refAmt
	}

	used, err := s.repo.HasUsedTicket(orderID)
	if err != nil {
		return nil, response.NewInternalServerError("failed to check user tickets", err)
	}
	if used {
		return nil, response.NewBadRequest("you have already used one of the tickets")
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

	s.repo.UpdatePaymentStatus(order.ID.String(), "refunded")
	s.repo.IncreaseUserBalance(userID, totalRefund)

	user, _ := s.user.GetUserByID(userID)
	return &dto.RefundOrderResponse{
		OrderID:      order.ID.String(),
		RefundAmount: totalRefund,
		RefundedAt:   now,
		UserBalance:  user.Balance,
	}, nil
}
