package repositories

import (
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	GetSummary() (*dto.SummaryReportResponse, error)
	GetAllUsers(params dto.UserQueryParams) ([]models.User, int64, error)
	GetAllEvents(params dto.EventQueryParams) ([]models.Event, int64, error)
	GetOrderReports(params dto.OrderReportQueryParams) ([]models.Order, int64, error)
	GetTicketSalesReports(params dto.TicketReportQueryParams) ([]models.Ticket, int64, error)
	GetPaymentReports(params dto.PaymentReportQueryParams) ([]models.Payment, int64, error)
	GetRefundReports(params dto.RefundReportQueryParams) ([]models.Order, int64, error)
	GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]models.WithdrawalRequest, int64, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

// TODO : move each function to its repository file for better organization (this for fast building purpose)
// ?NOTE : Raw SQL cmd instead of gorm methods for performance reasons
// repositories/report.go
func (r *adminRepository) GetSummary() (*dto.SummaryReportResponse, error) {
	var resp dto.SummaryReportResponse

	// Get Users Summary
	if err := r.db.Raw(`
		SELECT 
			COUNT(*) as total,
			COUNT(CASE WHEN created_at >= DATE_SUB(NOW(), INTERVAL 1 MONTH) THEN 1 END) as new_this_month,
			COUNT(CASE WHEN created_at >= DATE_SUB(NOW(), INTERVAL 30 DAY) THEN 1 END) as active_users
		FROM users
	`).Scan(&resp.Users).Error; err != nil {
		return nil, err
	}

	// Get Events Summary
	if err := r.db.Raw(`
		SELECT 
			COUNT(*) as total,
			COUNT(CASE WHEN status = 'active' THEN 1 END) as active,
			COUNT(CASE WHEN status = 'ongoing' THEN 1 END) as ongoing,
			COUNT(CASE WHEN status = 'done' THEN 1 END) as done,
			COUNT(CASE WHEN status = 'cancelled' THEN 1 END) as cancelled
		FROM events
	`).Scan(&resp.Events).Error; err != nil {
		return nil, err
	}

	// Get Orders Summary
	if err := r.db.Raw(`
		SELECT 
			COUNT(*) as total,
			COUNT(CASE WHEN status = 'pending' THEN 1 END) as pending,
			COUNT(CASE WHEN status = 'paid' THEN 1 END) as paid,
			COUNT(CASE WHEN status = 'failed' THEN 1 END) as failed,
			COUNT(CASE WHEN status = 'cancelled' THEN 1 END) as cancelled
		FROM orders
	`).Scan(&resp.Orders).Error; err != nil {
		return nil, err
	}

	// Get Revenue Summary
	if err := r.db.Raw(`
		SELECT 
			COALESCE(SUM(CASE WHEN status = 'paid' THEN total_price ELSE 0 END), 0) as total_revenue,
			COALESCE(SUM(CASE WHEN status = 'paid' AND created_at >= DATE_SUB(NOW(), INTERVAL 1 MONTH) THEN total_price ELSE 0 END), 0) as this_month,
			COALESCE(SUM(CASE WHEN status = 'pending' THEN total_price ELSE 0 END), 0) as pending_payments
		FROM orders
	`).Scan(&resp.Revenue).Error; err != nil {
		return nil, err
	}

	// Get Withdrawals Summary
	if err := r.db.Raw(`
		SELECT 
			COUNT(CASE WHEN status = 'pending' THEN 1 END) as pending,
			COALESCE(SUM(CASE WHEN status = 'pending' THEN amount ELSE 0 END), 0) as total_amount
		FROM withdrawal_requests
	`).Scan(&resp.Withdrawals).Error; err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *adminRepository) GetAllEvents(params dto.EventQueryParams) ([]models.Event, int64, error) {
	var events []models.Event
	var count int64

	// Get ALL events including those without tickets (for admin view)
	db := r.db.Model(&models.Event{})

	// Apply filters
	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("title LIKE ? OR description LIKE ?", like, like)
	}

	if params.Location != "" && params.Location != "all" {
		db = db.Where("location = ?", params.Location)
	}

	if params.Status != "" && params.Status != "all" {
		db = db.Where("status = ?", params.Status)
	}

	if params.StartDate != "" {
		db = db.Where("date >= ?", params.StartDate)
	}

	if params.EndDate != "" {
		db = db.Where("date <= ?", params.EndDate)
	}

	// Apply sorting
	switch params.Sort {
	case "date_asc":
		db = db.Order("date ASC")
	case "date_desc":
		db = db.Order("date DESC")
	case "title_asc":
		db = db.Order("title ASC")
	case "title_desc":
		db = db.Order("title DESC")
	case "created_asc":
		db = db.Order("created_at ASC")
	case "created_desc":
		db = db.Order("created_at DESC")
	default:
		db = db.Order("created_at DESC") // Default to newest first for admin
	}

	// Apply pagination defaults
	if params.Page <= 0 {
		params.Page = 1
	}

	if params.Limit <= 0 {
		params.Limit = 10
	}

	offset := (params.Page - 1) * params.Limit

	// Get total count
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Get events with tickets preloaded
	if err := db.Preload("Tickets").Limit(params.Limit).Offset(offset).Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, count, nil
}

func (r *adminRepository) GetAllUsers(params dto.UserQueryParams) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	db := r.db.Model(&models.User{})

	if params.Q != "" {
		db = db.Where("email LIKE ? OR fullname LIKE ?", "%"+params.Q+"%", "%"+params.Q+"%")
	}
	if params.Role != "" && params.Role != "all" {
		db = db.Where("role = ?", params.Role)
	}

	switch params.Sort {
	case "joined_asc":
		db = db.Order("created_at asc")
	case "joined_desc":
		db = db.Order("created_at desc")
	case "email_asc":
		db = db.Order("email asc")
	case "email_desc":
		db = db.Order("email desc")
	case "name_asc":
		db = db.Order("fullname asc")
	case "name_desc":
		db = db.Order("fullname desc")
	default:
		db = db.Order("created_at desc")
	}

	offset := (params.Page - 1) * params.Limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Limit(params.Limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func (r *adminRepository) GetOrderReports(params dto.OrderReportQueryParams) ([]models.Order, int64, error) {
	var orders []models.Order
	var count int64

	db := r.db.Model(&models.Order{}).Preload("Event")

	if params.Status != "" {
		db = db.Where("status = ?", params.Status)
	}

	if params.EventID != "" {
		db = db.Where("event_id = ?", params.EventID)
	}

	if params.DateFrom != "" {
		if fromDate, err := time.Parse("2006-01-02", params.DateFrom); err == nil {
			db = db.Where("created_at >= ?", fromDate)
		}
	}
	if params.DateTo != "" {
		if toDate, err := time.Parse("2006-01-02", params.DateTo); err == nil {
			toDate = toDate.Add(24 * time.Hour)
			db = db.Where("created_at < ?", toDate)
		}
	}

	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("fullname LIKE ? OR email LIKE ?", like, like)
	}

	db = db.Order("created_at DESC")

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (params.Page - 1) * params.Limit
	if err := db.Limit(params.Limit).Offset(offset).Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, count, nil
}

func (r *adminRepository) GetPaymentReports(params dto.PaymentReportQueryParams) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var count int64
	db := r.db.Model(&models.Payment{}).Joins("JOIN orders ON payments.order_id = orders.id")

	if params.Status != "" {
		db = db.Where("payments.status = ?", params.Status)
	}
	if params.Method != "" {
		db = db.Where("payments.method = ?", params.Method)
	}
	if params.Q != "" {
		q := "%" + params.Q + "%"
		db = db.Where("orders.fullname LIKE ? OR orders.email LIKE ?", q, q)
	}

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	offset := (params.Page - 1) * params.Limit
	if err := db.Preload("Order").
		Order("payments.created_at DESC").
		Limit(params.Limit).
		Offset(offset).
		Find(&payments).Error; err != nil {
		return nil, 0, err
	}

	return payments, count, nil
}

func (r *adminRepository) GetTicketSalesReports(params dto.TicketReportQueryParams) ([]models.Ticket, int64, error) {
	var tickets []models.Ticket
	var count int64

	db := r.db.Model(&models.Ticket{}).Joins("JOIN events ON tickets.event_id = events.id")

	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("events.title LIKE ? OR tickets.name LIKE ?", like, like)
	}

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (params.Page - 1) * params.Limit
	if err := db.Preload("Event").
		Order("events.created_at DESC").
		Limit(params.Limit).
		Offset(offset).
		Find(&tickets).Error; err != nil {
		return nil, 0, err
	}

	return tickets, count, nil
}

func (r *adminRepository) GetRefundReports(params dto.RefundReportQueryParams) ([]models.Order, int64, error) {
	var orders []models.Order
	var count int64
	db := r.db.Model(&models.Order{}).Preload("Event").Where("is_refunded = ?", true)

	if params.Q != "" {
		q := "%" + params.Q + "%"
		db = db.Where("fullname LIKE ? OR email LIKE ?", q, q)
	}

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	offset := (params.Page - 1) * params.Limit
	if err := db.Order("refunded_at DESC").
		Limit(params.Limit).
		Offset(offset).
		Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, count, nil
}

func (r *adminRepository) GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]models.WithdrawalRequest, int64, error) {
	var withdrawals []models.WithdrawalRequest
	var count int64

	db := r.db.Model(&models.WithdrawalRequest{}).
		Joins("JOIN users ON withdrawal_requests.user_id = users.id").
		Preload("User").
		Order("created_at DESC")

	if params.Q != "" {
		q := "%" + params.Q + "%"
		db = db.Where("users.fullname LIKE ? OR users.email LIKE ?", q, q)
	}

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (params.Page - 1) * params.Limit
	if err := db.Limit(params.Limit).Offset(offset).Find(&withdrawals).Error; err != nil {
		return nil, 0, err
	}

	return withdrawals, count, nil
}
