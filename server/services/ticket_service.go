package services

import (
	"server/dto"
	customErr "server/errors"
	"server/models"
	"server/repositories"

	"github.com/google/uuid"
)

type TicketService interface {
	DeleteTicket(id string) error
	CreateTicket(req dto.CreateTicketRequest, eventID string) error
	GetTicketByID(id string) (*dto.TicketResponse, error)
	UpdateTicket(id string, req dto.UpdateTicketRequest) error
}

type ticketService struct {
	repo  repositories.TicketRepository
	event repositories.EventRepository
}

func NewTicketService(repo repositories.TicketRepository, event repositories.EventRepository) TicketService {
	return &ticketService{repo, event}
}

func (s *ticketService) CreateTicket(req dto.CreateTicketRequest, eventID string) error {
	if req.Price < 0 || req.Quota < 0 || req.Limit < 0 {
		return customErr.NewBadRequest("invalid input: price, quota, or limit must not be negative")
	}

	event, err := s.event.GetEventByID(eventID)
	if err != nil || event == nil {
		return customErr.NewNotFound("event not found")
	}

	newTicket := &models.Ticket{
		ID:         uuid.New(),
		Name:       req.Name,
		EventID:    event.ID,
		Price:      req.Price,
		Limit:      req.Limit,
		Quota:      req.Quota,
		Sold:       0,
		Refundable: req.Refundable,
	}

	if err := s.repo.CreateTicket(newTicket); err != nil {
		return customErr.NewInternalServerError("failed to create ticket", err)
	}

	return nil
}

func (s *ticketService) GetTicketByID(id string) (*dto.TicketResponse, error) {
	ticket, err := s.repo.GetTicketByID(id)
	if err != nil || ticket == nil {
		return nil, customErr.NewNotFound("ticket not found").WithContext("ticketID", id)
	}

	return &dto.TicketResponse{
		ID:         ticket.ID.String(),
		EventID:    ticket.EventID.String(),
		Name:       ticket.Name,
		Price:      ticket.Price,
		Limit:      ticket.Limit,
		Quota:      ticket.Quota,
		Refundable: ticket.Refundable,
	}, nil
}

func (s *ticketService) DeleteTicket(id string) error {
	ticket, err := s.repo.GetTicketByID(id)
	if err != nil || ticket == nil {
		return customErr.NewNotFound("ticket not found").WithContext("ticketID", id)
	}

	if ticket.Sold > 0 {
		return customErr.NewBadRequest("cannot delete ticket that has been sold")
	}

	return s.repo.DeleteTicket(id)
}

func (s *ticketService) UpdateTicket(id string, req dto.UpdateTicketRequest) error {
	ticket, err := s.repo.GetTicketByID(id)
	if err != nil || ticket == nil {
		return customErr.NewNotFound("ticket not found").WithContext("ticketID", id)
	}

	if req.Price < 0 || req.Quota < 0 || req.Limit < 0 {
		return customErr.NewBadRequest("invalid input: price, quota, or limit must not be negative")
	}

	ticket.Name = req.Name
	ticket.Price = req.Price
	ticket.Limit = req.Limit
	ticket.Quota = req.Quota
	ticket.Refundable = req.Refundable

	return s.repo.UpdateTicket(ticket)
}
