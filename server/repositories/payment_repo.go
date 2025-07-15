package repositories

import (
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	ExpireOldPendingPayments() (int64, error)
	UpdatePayment(payment *models.Payment) error
	GetPaymentByID(paymentID string) (*models.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) GetPaymentByID(paymentID string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, "id = ?", paymentID).Error
	return &payment, err
}

func (r *paymentRepository) UpdatePayment(payment *models.Payment) error {
	return r.db.Model(&models.Payment{}).
		Where("id = ?", payment.ID).
		Updates(map[string]any{
			"method":  payment.Method,
			"status":  payment.Status,
			"paid_at": payment.PaidAt,
		}).Error
}

func (r *paymentRepository) ExpireOldPendingPayments() (int64, error) {
	threshold := time.Now().Add(-15 * time.Minute)

	var payments []models.Payment

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("status = ? AND created_at <= ?", "pending", threshold).Find(&payments).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Payment{}).
			Where("status = ? AND created_at <= ?", "pending", threshold).
			Update("status", "failed").Error; err != nil {
			return err
		}

		for _, p := range payments {
			if err := tx.Model(&models.Order{}).
				Where("id = ?", p.OrderID).
				Update("status", "failed").Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return int64(len(payments)), nil
}
