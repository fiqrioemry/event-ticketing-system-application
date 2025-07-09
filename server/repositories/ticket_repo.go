package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type TicketRepository interface {
	DeleteTicket(ID string) error
	CreateTicket(ticket *models.Ticket) error
	UpdateTicket(ticket *models.Ticket) error
	GetTicketByID(ID string) (*models.Ticket, error)
	GetTicketByEventID(eventID string) (*models.Ticket, error)
	GetAllTicketsByEventID(eventID string) ([]*models.Ticket, error)
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) CreateTicket(ticket *models.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) DeleteTicket(ID string) error {
	return r.db.Delete(&models.Ticket{}, "id = ?", ID).Error
}

func (r *ticketRepository) UpdateTicket(ticket *models.Ticket) error {
	return r.db.Save(ticket).Error
}

func (r *ticketRepository) GetTicketByID(ID string) (*models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.Where("id = ?", ID).First(&ticket).Error
	return &ticket, err
}

func (r *ticketRepository) GetTicketByEventID(eventID string) (*models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.Where("event_id = ?", eventID).First(&ticket).Error
	return &ticket, err
}

func (r *ticketRepository) GetAllTicketsByEventID(eventID string) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	err := r.db.Where("event_id = ?", eventID).Find(&tickets).Error
	return tickets, err
}
