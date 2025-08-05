package repositories

import (
	"errors"

	"github.com/fiqrioemry/event_ticketing_system_app/server/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(data *models.User) error
	UpdateUser(data *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
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
