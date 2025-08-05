package services

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/google/uuid"
)

type TicketService interface {
	DeleteTicket(id string) error
	GetTicketByID(id string) (*dto.TicketResponse, error)
	UpdateTicket(id string, req dto.UpdateTicketRequest) (*models.Ticket, error)
	CreateTicket(req dto.CreateTicketRequest, eventID string) (*models.Ticket, error)
}

type ticketService struct {
	repo  repositories.TicketRepository
	event repositories.EventRepository
}

func NewTicketService(repo repositories.TicketRepository, event repositories.EventRepository) TicketService {
	return &ticketService{repo, event}
}

func (s *ticketService) CreateTicket(req dto.CreateTicketRequest, eventID string) (*models.Ticket, error) {
	if req.Price < 0 || req.Quota < 0 || req.Limit < 0 {
		return nil, response.NewBadRequest("invalid input: price, quota, or limit must not be negative")
	}

	event, err := s.event.GetEventByID(eventID)
	if err != nil || event == nil {
		return nil, response.NewNotFound("event not found")
	}

	newTicket := &models.Ticket{
		ID:         uuid.New(),
		Name:       req.Name,
		EventID:    event.ID,
		Price:      req.Price,
		Limit:      req.Limit,
		Quota:      req.Quota,
		Refundable: req.Refundable,
	}
	if err := s.repo.CreateTicketWithEventStatusUpdate(newTicket, eventID); err != nil {
		return nil, response.NewInternalServerError("failed to create ticket", err)
	}

	return newTicket, nil
}

func (s *ticketService) GetTicketByID(id string) (*dto.TicketResponse, error) {
	ticket, err := s.repo.GetTicketByID(id)
	if err != nil || ticket == nil {
		return nil, response.NewNotFound("ticket not found")
	}

	ticketResponse := &dto.TicketResponse{
		ID:         ticket.ID.String(),
		EventID:    ticket.EventID.String(),
		Name:       ticket.Name,
		Price:      ticket.Price,
		Limit:      ticket.Limit,
		Quota:      ticket.Quota,
		Refundable: ticket.Refundable,
	}

	return ticketResponse, nil
}

func (s *ticketService) DeleteTicket(id string) error {
	ticket, err := s.repo.GetTicketByID(id)
	if err != nil || ticket == nil {
		return response.NewNotFound("ticket not found")
	}

	if ticket.Sold > 0 {
		return response.NewBadRequest("cannot delete ticket that has been sold")
	}

	return s.repo.DeleteTicket(id)
}

func (s *ticketService) UpdateTicket(id string, req dto.UpdateTicketRequest) (*models.Ticket, error) {
	ticket, err := s.repo.GetTicketByID(id)
	if err != nil || ticket == nil {
		return nil, response.NewNotFound("ticket not found")
	}

	if req.Price < 0 || req.Quota < 0 || req.Limit < 0 {
		return nil, response.NewBadRequest("invalid input: price, quota, or limit must not be negative")
	}

	ticket.Name = req.Name
	ticket.Price = req.Price
	ticket.Limit = req.Limit
	ticket.Quota = req.Quota
	ticket.Refundable = req.Refundable

	err = s.repo.UpdateTicket(ticket)
	if err != nil {
		return nil, response.NewInternalServerError("failed to update ticket", err)
	}

	return ticket, nil
}
