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
	GetAllUsers(params dto.UserQueryParams) ([]models.User, int64, error)
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
	return &user, err
}
func (r *userRepository) GetAllUsers(params dto.UserQueryParams) ([]models.User, int64, error) {
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
