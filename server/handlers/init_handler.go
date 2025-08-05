package handlers

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/services"
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
	AdminHandler      *AdminHandler
}

func InitHandlers(s *services.Services, r *repositories.Repositories) *Handlers {
	return &Handlers{
		AuthHandler:       NewAuthHandler(s.AuthService),
		OrderHandler:      NewOrderHandler(s.OrderService),
		UserTicketHandler: NewUserTicketHandler(s.UserTicketService),
		UserHandler:       NewUserHandler(s.UserService, r.AuditRepository),
		EventHandler:      NewEventHandler(s.EventService, r.AuditRepository),
		TicketHandler:     NewTicketHandler(s.TicketService, r.AuditRepository),
		WithdrawalHandler: NewWithdrawalHandler(s.WithdrawalService, r.AuditRepository),
		PaymentHandler:    NewPaymentHandler(s.PaymentService),
		AdminHandler:      NewAdminHandler(s.AdminService),
	}
}
