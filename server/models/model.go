package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:text;not null"`
	Balance   float64   `gorm:"type:decimal(12,2);default:0.00"`
	AvatarURL string    `gorm:"type:varchar(255);default:''"`
	Role      string    `gorm:"type:enum('admin','user');default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Events    []Event    `gorm:"foreignKey:CreatedBy"`
	Purchases []Purchase `gorm:"foreignKey:UserID"`
}

type Event struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Title       string    `gorm:"type:varchar(150);unique;not null"`
	Description string    `gorm:"type:text"`
	Location    string    `gorm:"type:varchar(100)"`
	StartTime   time.Time
	EndTime     time.Time
	Status      string    `gorm:"type:enum('active','ongoing','done','cancelled');default:'active'"`
	CreatedBy   uuid.UUID `gorm:"type:char(36)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Categories []TicketCategory `gorm:"foreignKey:EventID"`
	Tickets    []Ticket         `gorm:"foreignKey:EventID"`
	Reports    []Report         `gorm:"foreignKey:EventID"`
}

type TicketCategory struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	EventID   uuid.UUID `gorm:"type:char(36);index"`
	Name      string    `gorm:"type:varchar(50);not null"` // VIP, Reguler, Early Bird
	Price     float64   `gorm:"type:decimal(12,2);not null"`
	Quota     int       `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Event   Event    `gorm:"foreignKey:EventID"`
	Tickets []Ticket `gorm:"foreignKey:CategoryID"`
}

type Ticket struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	EventID    uuid.UUID `gorm:"type:char(36);index"`
	CategoryID uuid.UUID `gorm:"type:char(36);index"`
	SeatNo     string    `gorm:"type:varchar(10)"`
	Status     string    `gorm:"type:enum('available','booked','cancelled');default:'available'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Event    Event          `gorm:"foreignKey:EventID"`
	Category TicketCategory `gorm:"foreignKey:CategoryID"`
	Purchase *Purchase      `gorm:"foreignKey:TicketID"`
}

type Purchase struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID       uuid.UUID `gorm:"type:char(36);index"`
	TicketID     uuid.UUID `gorm:"type:char(36);uniqueIndex"`
	Status       string    `gorm:"type:enum('pending','completed','cancelled');default:'pending'"`
	PurchaseTime time.Time
	PaymentID    *uuid.UUID `gorm:"type:char(36);"`

	User    User
	Ticket  Ticket
	Payment *Payment
}

type Payment struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Method    string    `gorm:"type:varchar(50)"`
	Amount    float64   `gorm:"type:decimal(12,2)"`
	Status    string    `gorm:"type:enum('pending','paid','failed','refunded');default:'pending'"`
	PaidAt    *time.Time
	CreatedAt time.Time

	Purchases []Purchase `gorm:"foreignKey:PaymentID"`
}

type Report struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey"`
	EventID      uuid.UUID `gorm:"type:char(36);index"`
	TotalSold    int       `gorm:"not null"`
	TotalRevenue float64   `gorm:"type:decimal(12,2);not null"`
	GeneratedAt  time.Time

	Event Event `gorm:"foreignKey:EventID"`
}
