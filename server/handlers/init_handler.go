package handlers

import (
	"server/services"
)

type Handlers struct {
	AuthHandler       *AuthHandler
	UserHandler       *UserHandler
	EventHandler      *EventHandler
	TicketHandler     *TicketHandler
	OrderHandler      *OrderHandler
	UserTicketHandler *UserTicketHandler
	WithdrawalHandler *WithdrawalHandler
	PaymentHandler    *PaymentHandler
	ReportHandler     *ReportHandler
}

func InitHandlers(s *services.Services) *Handlers {
	return &Handlers{
		AuthHandler:       NewAuthHandler(s.AuthService),
		UserHandler:       NewUserHandler(s.UserService),
		EventHandler:      NewEventHandler(s.EventService),
		TicketHandler:     NewTicketHandler(s.TicketService),
		OrderHandler:      NewOrderHandler(s.OrderService),
		UserTicketHandler: NewUserTicketHandler(s.UserTicketService),
		WithdrawalHandler: NewWithdrawalHandler(s.WithdrawalService),
		PaymentHandler:    NewPaymentHandler(s.PaymentService),
		ReportHandler:     NewReportHandler(s.ReportService),
	}
}
