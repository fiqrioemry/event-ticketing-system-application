package models

import (
	"time"

	"github.com/google/uuid"
)

// ? Separate each models into it own files
// TODO : sepearate based on modules if project is getting bigger

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Fullname  string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:text;not null"`
	AvatarURL string    `gorm:"type:varchar(255);default:''"`
	Role      string    `gorm:"type:enum('admin','user');default:'user'"`
	Balance   float64   `gorm:"type:decimal(12,2);default:0.00"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// TODO : Future improvements add categories and tags for events, so event can be classfied and filtered based on these attributes
//
//	type Category struct {
//		ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
//		Name      string    `gorm:"type:varchar(100);unique;not null"`
//		CreatedAt time.Time `gorm:"autoCreateTime"`
//	}
//

//  CategoryID uuid.UUID `gorm:"type:char(36);index"` // TODO : Future improvements, add category for event
// TODO : Add Slug for better SEO and URL structure

type Event struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Image       string    `gorm:"type:varchar(255);default:''"`
	Title       string    `gorm:"type:varchar(150);unique;not null"`
	Description string    `gorm:"type:text"`
	Location    string    `gorm:"type:varchar(100)"`
	Date        time.Time `gorm:"not null" json:"date"`
	StartTime   int       `gorm:"not null" json:"startTime"`
	EndTime     int       `gorm:"not null" json:"endTime"`
	Status      string    `gorm:"type:enum('active','ongoing','done','cancelled');default:'active'"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	// Category category `gorm:"foreignKey:CategoryID"` // TODO : Future improvements, add category for event
	Tickets []Ticket `gorm:"foreignKey:EventID"`
}

type Ticket struct {
	ID            uuid.UUID `gorm:"type:char(36);primaryKey"`
	EventID       uuid.UUID `gorm:"type:char(36);index"`
	Name          string    `gorm:"type:varchar(100);not null"`
	Price         float64   `gorm:"type:decimal(12,2);not null"`
	Limit         int       `gorm:"not null"`
	Quota         int       `gorm:"not null"`
	Sold          int       `gorm:"default:0"`
	Refundable    bool      `gorm:"default:false"`
	RefundPercent int       `gorm:"type:int;default:50"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`

	Event Event `gorm:"foreignKey:EventID"`
}

type Order struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID     uuid.UUID `gorm:"type:char(36);index"`
	EventID    uuid.UUID `gorm:"type:char(36);index"`
	Fullname   string    `gorm:"type:varchar(100);not null"`
	Email      string    `gorm:"type:varchar(100);not null"`
	Phone      string    `gorm:"type:varchar(20);not null"`
	TotalPrice float64   `gorm:"type:decimal(12,2);not null"`
	Status     string    `gorm:"type:enum('pending','paid','failed','cancelled','refunded');default:'pending'"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	// TODO : Optional features. if sufficient time, implement these
	IsRefunded   bool       `gorm:"default:false"`
	RefundedAt   *time.Time `gorm:"default:null"`
	RefundAmount float64    `gorm:"type:decimal(12,2);default:0"`
	RefundReason string     `gorm:"type:text;default:null"`

	Event Event `gorm:"foreignKey:EventID"`
}

type OrderDetail struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	OrderID    uuid.UUID `gorm:"type:char(36);index"`
	TicketID   uuid.UUID `gorm:"type:char(36);index"`
	TicketName string    `gorm:"type:varchar(100);not null"`
	Quantity   int       `gorm:"not null"`
	Price      float64   `gorm:"type:decimal(12,2);not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

type Payment struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID  `gorm:"type:char(36);index"`
	OrderID   uuid.UUID  `gorm:"type:char(36);index"`
	Fullname  string     `gorm:"type:varchar(100);not null"`
	Email     string     `gorm:"type:varchar(100);not null"`
	Method    string     `gorm:"type:varchar(50)"`
	Amount    float64    `gorm:"type:decimal(12,2)"`
	Status    string     `gorm:"type:enum('pending','paid','failed','refunded');default:'pending'"`
	PaidAt    *time.Time `gorm:"default:null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`

	Order Order `gorm:"foreignKey:OrderID"`
}

type UserTicket struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID `gorm:"type:char(36);index"`
	EventID   uuid.UUID `gorm:"type:char(36);index"`
	TicketID  uuid.UUID `gorm:"type:char(36);index"`
	IsUsed    bool      `gorm:"default:false"`
	UsedAt    *time.Time
	QRCode    string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Ticket Ticket `gorm:"foreignKey:TicketID"`
	Event  Event  `gorm:"foreignKey:EventID"`
}

type WithdrawalRequest struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID     uuid.UUID `gorm:"type:char(36);index"`
	Amount     float64   `gorm:"type:decimal(12,2);not null"`
	Status     string    `gorm:"type:enum('pending','approved','rejected');default:'pending'"`
	Reason     string    `gorm:"type:text"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	ApprovedAt *time.Time

	User User `gorm:"foreignKey:UserID"`
}
