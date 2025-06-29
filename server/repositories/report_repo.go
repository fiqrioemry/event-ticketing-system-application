package repositories

import (
	"server/dto"
	"server/models"
	"time"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetSummary() (*dto.SummaryReportResponse, error)
	GetOrderReports(params dto.OrderReportQueryParams) ([]models.Order, int64, error)
	GetTicketSalesReports(params dto.TicketReportQueryParams) ([]models.Ticket, int64, error)
	GetPaymentReports(params dto.PaymentReportQueryParams) ([]models.Payment, int64, error)
	GetRefundReports(params dto.RefundReportQueryParams) ([]models.Order, int64, error)
	GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]models.WithdrawalRequest, int64, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db}
}

// TODO : move each function to its repository file for better organization (this for fast building purpose)
// ?NOTE : Raw SQL cmd instead of gorm methods for performance reasons
func (r *reportRepository) GetSummary() (*dto.SummaryReportResponse, error) {
	var resp dto.SummaryReportResponse
	if err := r.db.Raw(`
		SELECT 
			COALESCE(SUM(total_price),0) as total_revenue,
			COUNT(*) as total_orders,
			(SELECT COUNT(*) FROM users) as total_users,
			(SELECT COUNT(*) FROM events) as total_events,
			COALESCE(SUM(refund_amount),0) as total_refund,
			(SELECT SUM(quantity) FROM order_details) as total_ticket_sold
		FROM orders
		WHERE status='paid'
	`).Scan(&resp).Error; err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *reportRepository) GetOrderReports(params dto.OrderReportQueryParams) ([]models.Order, int64, error) {
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

func (r *reportRepository) GetPaymentReports(params dto.PaymentReportQueryParams) ([]models.Payment, int64, error) {
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

func (r *reportRepository) GetTicketSalesReports(params dto.TicketReportQueryParams) ([]models.Ticket, int64, error) {
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

func (r *reportRepository) GetRefundReports(params dto.RefundReportQueryParams) ([]models.Order, int64, error) {
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

func (r *reportRepository) GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]models.WithdrawalRequest, int64, error) {
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
