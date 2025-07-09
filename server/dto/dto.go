package dto

import (
	"mime/multipart"
	"time"
)

// ? Separate each dto into sections based on their functionality
// TODO : sepearate based on modules if project is getting bigger

// 1. USER & AUTHENTICATION MODULE MANAGEMENT =============
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Fullname string `json:"fullname" binding:"required,min=5"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	User         ProfileResponse `json:"user"`
	AccessToken  string          `json:"accessToken"`
	RefreshToken string          `json:"refreshToken"`
}

type ResendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}

type ProfileResponse struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Fullname string    `json:"fullname"`
	Avatar   string    `json:"avatar"`
	Role     string    `json:"role"`
	Balance  float64   `json:"balance"`
	JoinedAt time.Time `json:"joinedAt"`
}

type UpdateProfileRequest struct {
	Fullname  string                `form:"fullname" binding:"required,min=5"`
	Avatar    *multipart.FileHeader `form:"avatar" binding:"omitempty"`
	AvatarURL string                `form:"avatarURL"`
}

type UserQueryParams struct {
	Q     string `form:"q"`
	Role  string `form:"role"`
	Sort  string `form:"sort"`
	Page  int    `form:"page,default=1"`
	Limit int    `form:"limit,default=10"`
}

type UserListResponse struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Fullname string    `json:"fullname"`
	Avatar   string    `json:"avatar"`
	JoinedAt time.Time `json:"joinedAt"`
}

type UserDetailResponse struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Fullname string    `json:"fullname"`
	Avatar   string    `json:"avatar"`
	JoinedAt time.Time `json:"joinedAt"`
}

// 2. EVENT MODULE MANAGEMENT =============
type EventQueryParams struct {
	Q         string `form:"search"`
	Status    string `form:"status"`
	StartDate string `form:"startDate"`
	EndDate   string `form:"endDate"`
	Location  string `form:"location"`
	Sort      string `form:"sort"`
	Page      int    `form:"page" default:"1"`
	Limit     int    `form:"limit" default:"10"`
}

type EventResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	IsAvailable bool      `json:"isAvailable"`
	StartPrice  float64   `json:"startPrice"`
	StartTime   int       `json:"startTime"`
	EndTime     int       `json:"endTime"`
	Date        time.Time `json:"date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

type EventDetailResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	StartPrice  float64   `json:"startPrice"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	StartTime   int       `json:"startTime"`
	Date        time.Time `json:"date"`
	EndTime     int       `json:"endTime"`
	CreatedAt   time.Time `json:"createdAt"`

	Tickets []TicketResponse `json:"tickets"`
}

type UpdateEventRequest struct {
	Title       string    `form:"title" binding:"required,min=5,max=150"`
	Description string    `form:"description" binding:"required"`
	Location    string    `form:"location" binding:"required"`
	Date        time.Time `form:"date" binding:"required"`
	StartTime   int       `form:"startTime" binding:"required,min=0,max=23"`
	EndTime     int       `form:"endTime" binding:"required,min=1,max=24"`
	Status      string    `form:"status" binding:"required,oneof=active ongoing done cancelled"`

	Image    *multipart.FileHeader `form:"image"`
	ImageURL string                `form:"-"`
}

type TicketInfo struct {
	ID     string `json:"id"`
	Status string `json:"status"` // available, booked, cancelled
}

type CreateEventRequest struct {
	Title       string                `json:"title" binding:"required,min=5,max=150"`
	Description string                `json:"description" binding:"required"`
	Location    string                `json:"location" binding:"required"`
	Date        time.Time             `json:"date" binding:"required"`
	StartTime   int                   `json:"startTime" binding:"required,min=0,max=23"`
	EndTime     int                   `json:"endTime" binding:"required,min=1,max=24"`
	Status      string                `json:"status" binding:"required,oneof=active ongoing done cancelled"`
	Tickets     []CreateTicketRequest `json:"tickets" binding:"required,min=1,dive"`

	// For file upload (tidak di-unmarshal dari JSON)
	Image    *multipart.FileHeader `json:"-"`
	ImageURL string                `json:"-"`
}

type CreateTicketRequest struct {
	Name          string  `json:"name" binding:"required,min=3,max=50"`
	Price         float64 `json:"price" binding:"required,min=0"`
	Quota         int     `json:"quota" binding:"required,min=1"`
	Limit         int     `json:"limit" binding:"omitempty,min=1"`
	Refundable    bool    `json:"isRefundable"`
	RefundPercent int     `json:"refundPercent" binding:"omitempty,min=0,max=100"`
}

// 3. TICKET  MODULE MANAGEMENT =============
type TicketResponse struct {
	ID         string  `json:"id"`
	EventID    string  `json:"eventId"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quota      int     `json:"quota"`
	Limit      int     `json:"limit"`
	Sold       int     `json:"sold"`
	Refundable bool    `json:"isRefundable"`
}

type TicketQueryParams struct {
	EventID string `form:"event_id"`
	Status  string `form:"status"` // available, booked, cancelled
	Page    int    `form:"page" default:"1"`
	Limit   int    `form:"limit" default:"10"`
}

type UpdateTicketRequest struct {
	Name       string  `json:"name" binding:"required,min=3,max=50"`
	EventID    string  `json:"eventId" binding:"required,uuid"`
	Price      float64 `json:"price" binding:"required,min=0"`
	Quota      int     `json:"quota" binding:"required,min=1"`
	Limit      int     `json:"limit" binding:"omitempty,min=1"`
	Refundable bool    `json:"isRefundable" default:"false"`
}

// 4. ORDER MODULE MANAGEMENT =============
type CreateOrderRequest struct {
	EventID      string               `json:"eventId" binding:"required,uuid"`
	OrderDetails []OrderDetailRequest `json:"orderDetails" binding:"required,dive"`
	Fullname     string               `json:"fullname" binding:"required,min=3,max=100"`
	Email        string               `json:"email" binding:"required,email"`
	Phone        string               `json:"phone" binding:"required,min=10,max=15"`
}

type OrderDetailRequest struct {
	TicketID string `json:"ticketId" binding:"required,uuid"`
	Quantity int    `json:"quantity" binding:"required,min=1"`
}

type CheckoutSessionResponse struct {
	PaymentID string
	SessionID string
	URL       string
}

type OrderQueryParams struct {
	Q      string `form:"search"`
	Status string `form:"status"`
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
	Sort   string `form:"sort"`
}

type OrderResponse struct {
	ID         string    `json:"id"`
	EventID    string    `json:"eventId"`
	EventName  string    `json:"eventName"`
	EventImage string    `json:"eventImage"`
	Fullname   string    `json:"fullname"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	TotalPrice float64   `json:"totalPrice"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
}

type OrderDetailResponse struct {
	ID         string    `json:"id"`
	TicketName string    `json:"ticketName"`
	TicketID   string    `json:"ticketId"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"createdAt"`
}

// 5. USER TICKET MODULE MANAGEMENT =============
type UserTicketResponse struct {
	ID         string     `json:"id"`
	EventID    string     `json:"eventId"`
	TicketID   string     `json:"ticketId"`
	TicketName string     `json:"ticketName"`
	EventName  string     `json:"eventName"`
	QRCode     string     `json:"qrCode"`
	IsUsed     bool       `json:"isUsed"`
	UsedAt     *time.Time `json:"usedAt,omitempty"`
}

type ValidateTicketRequest struct {
	QRCode string `json:"qrCode" binding:"required"`
}

// 6. REFUND MODULE MANAGEMENT =============
type RefundOrderRequest struct {
	Reason string `json:"reason" binding:"required"`
}

type RefundOrderResponse struct {
	OrderID      string    `json:"orderId"`
	RefundAmount float64   `json:"refundAmount"`
	RefundedAt   time.Time `json:"refundedAt"`
	UserBalance  float64   `json:"userBalance"`
}

// 7. WITHDRAWAL MODULE MANAGEMENT =============

type CreateWithdrawalRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
	Reason string  `json:"reason" binding:"required"`
}

type WithdrawalResponse struct {
	ID         string     `json:"id"`
	UserID     string     `json:"userId"`
	Amount     float64    `json:"amount"`
	Status     string     `json:"status"`
	Reason     string     `json:"reason"`
	CreatedAt  time.Time  `json:"createdAt"`
	ReviewedBy string     `json:"reviewedBy,omitempty"`
	ApprovedAt *time.Time `json:"approvedAt,omitempty"`
}

// 8. REPORT MODULE MANAGEMENT =============

type SummaryReportResponse struct {
	TotalRevenue    float64 `json:"totalRevenue"`
	TotalOrders     int64   `json:"totalOrders"`
	TotalTicketSold int64   `json:"totalTicketSold"`
	TotalRefund     float64 `json:"totalRefund"`
	TotalUsers      int64   `json:"totalUsers"`
	TotalEvents     int64   `json:"totalEvents"`
}

type OrderReportQueryParams struct {
	Q        string `form:"search"`
	Page     int    `form:"page,default=1"`
	Limit    int    `form:"limit,default=10"`
	Status   string `form:"status"`
	EventID  string `form:"eventId"`
	DateFrom string `form:"dateFrom"`
	DateTo   string `form:"dateTo"`
	Export   string `form:"export" binding:"omitempty,oneof=csv pdf"`
}

type OrderReportResponse struct {
	OrderID    string    `json:"orderId"`
	Fullname   string    `json:"fullname"`
	Email      string    `json:"email"`
	EventTitle string    `json:"eventTitle"`
	TotalPrice float64   `json:"totalPrice"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
}

type TicketReportQueryParams struct {
	Q      string `form:"search"`
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
	Export string `form:"export" binding:"omitempty,oneof=csv pdf"`
}

type TicketSalesReportResponse struct {
	EventTitle  string  `json:"eventTitle"`
	TicketName  string  `json:"ticketName"`
	TicketPrice float64 `json:"ticketPrice"`
	Quota       int     `json:"quota"`
	Sold        int     `json:"sold"`
	Remaining   int     `json:"remaining"`
}

// payment report response
type PaymentReportQueryParams struct {
	Q      string `form:"search"`
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
	Status string `form:"status"`
	Method string `form:"method"`
	Export string `form:"export" binding:"omitempty,oneof=csv pdf"`
}

type PaymentReportResponse struct {
	PaymentID string     `json:"paymentId"`
	OrderID   string     `json:"orderId"`
	Fullname  string     `json:"fullname"`
	Email     string     `json:"email"`
	Method    string     `json:"method"`
	Amount    float64    `json:"amount"`
	Status    string     `json:"status"`
	PaidAt    *time.Time `json:"paidAt,omitempty"`
}

// refund report response
type RefundReportQueryParams struct {
	Q      string `form:"search"`
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
	Export string `form:"export" binding:"omitempty,oneof=csv pdf"`
}

type RefundReportResponse struct {
	OrderID      string     `json:"orderId"`
	Fullname     string     `json:"fullname"`
	Email        string     `json:"email"`
	EventTitle   string     `json:"eventTitle"`
	RefundAmount float64    `json:"refundAmount"`
	RefundReason string     `json:"refundReason"`
	RefundedAt   *time.Time `json:"refundedAt,omitempty"`
}

// Withdrawal
type WithdrawalReportQueryParams struct {
	Q      string `form:"search"`
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
	Export string `form:"export" binding:"omitempty,oneof=csv pdf"`
}

type WithdrawalReportResponse struct {
	WithdrawalID string     `json:"withdrawalId"`
	UserID       string     `json:"userId"`
	Fullname     string     `json:"fullname"`
	Email        string     `json:"email"`
	Amount       float64    `json:"amount"`
	Status       string     `json:"status"`
	Reason       string     `json:"reason"`
	CreatedAt    time.Time  `json:"createdAt"`
	ApprovedAt   *time.Time `json:"approvedAt,omitempty"`
}

// PAGINATION RESPONSE
type PaginationResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalRows  int `json:"totalRows"`
	TotalPages int `json:"totalPages"`
}
