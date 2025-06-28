package repositories

import (
	"server/dto"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetAllTransactions() ([]dto.TransactionReportResponse, error)
	GetTransactionByID(id string) (*dto.TransactionReportResponse, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db}
}

func (r *reportRepository) GetAllTransactions() ([]dto.TransactionReportResponse, error) {
	var rows []dto.TransactionReportResponse
	err := r.db.Raw(`
		SELECT 
			orders.id AS order_id,
			users.fullname AS user_name,
			users.email AS user_email,
			events.title AS event_title,
			payments.amount AS total_paid,
			orders.refund_amount,
			orders.status,
			payments.method,
			payments.paid_at,
			orders.refunded_at
		FROM orders
		JOIN users ON users.id = orders.user_id
		JOIN events ON events.id = orders.event_id
		JOIN payments ON payments.order_id = orders.id
	`).Scan(&rows).Error
	return rows, err
}

func (r *reportRepository) GetTransactionByID(id string) (*dto.TransactionReportResponse, error) {
	var row dto.TransactionReportResponse
	err := r.db.Raw(`
		SELECT 
			orders.id AS order_id,
			users.fullname AS user_name,
			users.email AS user_email,
			events.title AS event_title,
			payments.amount AS total_paid,
			orders.refund_amount,
			orders.status,
			payments.method,
			payments.paid_at,
			orders.refunded_at
		FROM orders
		JOIN users ON users.id = orders.user_id
		JOIN events ON events.id = orders.event_id
		JOIN payments ON payments.order_id = orders.id
		WHERE orders.id = ?
	`, id).Scan(&row).Error

	if err != nil {
		return nil, err
	}
	return &row, nil
}
