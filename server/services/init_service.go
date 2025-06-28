package services

import (
	"server/repositories"
)

type Services struct {
	UserService       UserService
	AuthService       AuthService
	EventService      EventService
	TicketService     TicketService
	OrderService      OrderService
	PaymentService    PaymentService
	UserTicketService UserTicketService
	WithdrawalService WithdrawalService
	ReportService     ReportService
}

func InitServices(r *repositories.Repositories) *Services {
	return &Services{
		UserService:       NewUserService(r.UserRepository),
		AuthService:       NewAuthService(r.AuthRepository),
		EventService:      NewEventService(r.EventRepository, r.TicketRepository),
		TicketService:     NewTicketService(r.TicketRepository),
		OrderService:      NewOrderService(r.OrderRepository, r.UserRepository, r.TicketRepository, r.EventRepository, r.UserTicketRepository),
		PaymentService:    NewPaymentService(r.PaymentRepository, r.OrderRepository, r.TicketRepository, r.UserTicketRepository),
		UserTicketService: NewUserTicketService(r.UserTicketRepository),
		WithdrawalService: NewWithdrawalService(r.WithdrawalRepository),
		ReportService:     NewReportService(r.ReportRepository),
	}
}
