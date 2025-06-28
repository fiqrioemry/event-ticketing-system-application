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
	CreateTicket(req dto.CreateTicketRequest) error
	GetTicketByID(id string) (*dto.TicketResponse, error)
	UpdateTicket(id string, req dto.UpdateTicketRequest) error
}

type ticketService struct {
	repo repositories.TicketRepository
}

func NewTicketService(repo repositories.TicketRepository) TicketService {
	return &ticketService{repo}
}

func (s *ticketService) CreateTicket(req dto.CreateTicketRequest) error {
	if req.Price < 0 || req.Quota < 0 || req.Limit < 0 {
		return customErr.NewBadRequest("invalid input: price, quota, or limit must not be negative")
	}

	ticket, err := s.repo.GetTicketByEventID(req.EventID)
	if err != nil {
		return customErr.NewInternal("failed to check existing ticket", err)
	}

	newTicket := &models.Ticket{
		ID:         uuid.New(),
		Name:       req.Name,
		EventID:    ticket.EventID,
		Price:      req.Price,
		Limit:      req.Limit,
		Quota:      req.Quota,
		Sold:       0,
		Refundable: req.Refundable,
	}

	if err := s.repo.CreateTicket(newTicket); err != nil {
		return customErr.NewInternal("failed to create ticket", err)
	}

	return nil
}

func (s *ticketService) GetTicketByID(id string) (*dto.TicketResponse, error) {
	ticket, err := s.repo.GetTicketByID(id)
	if err != nil {
		return nil, customErr.NewNotFound("ticket not found")
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
	if err != nil {
		return customErr.NewNotFound("ticket not found")
	}

	if ticket.Sold > 0 {
		return customErr.NewBadRequest("cannot delete ticket that has been sold")
	}

	return s.repo.DeleteTicket(id)
}

func (s *ticketService) UpdateTicket(id string, req dto.UpdateTicketRequest) error {
	ticket, err := s.repo.GetTicketByID(id)
	if err != nil {
		return customErr.NewNotFound("ticket not found")
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
