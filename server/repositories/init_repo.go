package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository       UserRepository
	AuthRepository       UserRepository
	EventRepository      EventRepository
	TicketRepository     TicketRepository
	UserTicketRepository UserTicketRepository
	OrderRepository      OrderRepository
	WithdrawalRepository WithdrawalRepository
	PaymentRepository    PaymentRepository
	ReportRepository     ReportRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:       NewUserRepository(db),
		AuthRepository:       NewUserRepository(db),
		EventRepository:      NewEventRepository(db),
		TicketRepository:     NewTicketRepository(db),
		UserTicketRepository: NewUserTicketRepository(db),
		OrderRepository:      NewOrderRepository(db),
		WithdrawalRepository: NewWithdrawalRepository(db),
		PaymentRepository:    NewPaymentRepository(db),
		ReportRepository:     NewReportRepository(db),
	}
}
