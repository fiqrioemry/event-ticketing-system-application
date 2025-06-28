package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type WithdrawalRepository interface {
	CreateWithdrawal(w *models.WithdrawalRequest) error
	GetAllWithdrawals() ([]models.WithdrawalRequest, error)
	GetWithdrawalByID(id string) (*models.WithdrawalRequest, error)
	UpdateWithdrawal(w *models.WithdrawalRequest) error
	GetUserByID(userID string) (*models.User, error)
	DecreaseUserBalance(userID string, amount float64) error
}

type withdrawalRepository struct {
	db *gorm.DB
}

func NewWithdrawalRepository(db *gorm.DB) WithdrawalRepository {
	return &withdrawalRepository{db}
}

func (r *withdrawalRepository) CreateWithdrawal(w *models.WithdrawalRequest) error {
	return r.db.Create(w).Error
}

func (r *withdrawalRepository) GetAllWithdrawals() ([]models.WithdrawalRequest, error) {
	var list []models.WithdrawalRequest
	err := r.db.Find(&list).Error
	return list, err
}

func (r *withdrawalRepository) GetWithdrawalByID(id string) (*models.WithdrawalRequest, error) {
	var w models.WithdrawalRequest
	err := r.db.First(&w, "id = ?", id).Error
	return &w, err
}

func (r *withdrawalRepository) UpdateWithdrawal(w *models.WithdrawalRequest) error {
	return r.db.Save(w).Error
}

func (r *withdrawalRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", userID).Error
	return &user, err
}

func (r *withdrawalRepository) DecreaseUserBalance(userID string, amount float64) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("balance", gorm.Expr("balance - ?", amount)).Error
}
