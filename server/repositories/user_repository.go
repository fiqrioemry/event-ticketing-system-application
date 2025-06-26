package repositories

import (
	"errors"
	"server/dto"
	"server/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(data *models.User) error
	UpdateUser(data *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetAllUser(params dto.UserQueryParams) ([]models.User, int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(data *models.User) error {
	return r.db.Create(data).Error
}

func (r *userRepository) UpdateUser(data *models.User) error {
	return r.db.Save(data).Error
}

func (r *userRepository) GetUserByID(id string) (*models.User, error) {
	var data models.User
	err := r.db.First(&data, "id = ?", id).Error
	return &data, err
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, nil
}

func (r *userRepository) GetAllUser(params dto.UserQueryParams) ([]models.User, int64, error) {
	var list []models.User
	var count int64

	db := r.db.Model(&models.User{})

	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("name LIKE ?", like)
	}

	db.Count(&count)

	offset := (params.Page - 1) * params.Limit
	err := db.Limit(params.Limit).Offset(offset).Order("created_at desc").Find(&list).Error

	return list, count, err
}
