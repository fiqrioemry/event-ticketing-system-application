package repositories

import (
	"errors"
	"server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTicketRepository interface {
	MarkTicketUsed(id string) error
	CreateUserTicket(ticket *models.UserTicket) error
	ValidateQRCode(qr string) (*models.UserTicket, error)
	GetUserTicketByID(id string) (*models.UserTicket, error)
	GetUserTicketsByOrderID(orderID string) ([]models.UserTicket, error)
	CountUserTicketsByTicketID(ticketID uuid.UUID) (int64, error)
}

type userTicketRepository struct {
	db *gorm.DB
}

func NewUserTicketRepository(db *gorm.DB) UserTicketRepository {
	return &userTicketRepository{db}
}

func (r *userTicketRepository) CreateUserTicket(ticket *models.UserTicket) error {
	return r.db.Create(ticket).Error
}
func (r *userTicketRepository) GetUserTicketsByOrderID(orderID string) ([]models.UserTicket, error) {
	var userTickets []models.UserTicket

	err := r.db.
		Model(&models.UserTicket{}).
		Select("user_tickets.*").
		Joins("JOIN order_details ON user_tickets.ticket_id = order_details.ticket_id").
		Where("order_details.order_id = ?", orderID).
		Group("user_tickets.id").
		Preload("Event").
		Preload("Ticket").
		Find(&userTickets).Error

	return userTickets, err
}

func (r *userTicketRepository) GetUserTicketByID(id string) (*models.UserTicket, error) {
	var ticket models.UserTicket
	err := r.db.Preload("Ticket").First(&ticket, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ticket, err
}

func (r *userTicketRepository) ValidateQRCode(qr string) (*models.UserTicket, error) {
	var ticket models.UserTicket
	err := r.db.Preload("Ticket").Preload("Event").Where("qr_code = ?", qr).First(&ticket).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ticket, err
}

func (r *userTicketRepository) MarkTicketUsed(id string) error {
	return r.db.Model(&models.UserTicket{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_used": true,
		"used_at": gorm.Expr("NOW()"),
	}).Error
}

func (r *userTicketRepository) CountUserTicketsByTicketID(ticketID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&models.UserTicket{}).Where("ticket_id = ?", ticketID).Count(&count).Error
	return count, err
}
