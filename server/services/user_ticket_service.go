package services

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/go-api-toolkit/response"
)

type UserTicketService interface {
	GetUserTicketByID(id string) (*dto.UserTicketResponse, error)
	ValidateTicket(qr string) (*dto.UserTicketResponse, error)
	MarkTicketUsed(id string) error
}

type userTicketService struct {
	repo repositories.UserTicketRepository
}

func NewUserTicketService(repo repositories.UserTicketRepository) UserTicketService {
	return &userTicketService{repo}
}

func (s *userTicketService) GetUserTicketByID(id string) (*dto.UserTicketResponse, error) {
	ticket, err := s.repo.GetUserTicketByID(id)
	if err != nil || ticket == nil {
		return nil, response.NewNotFound("ticket not found")
	}

	resp := &dto.UserTicketResponse{
		ID:         ticket.ID.String(),
		EventID:    ticket.EventID.String(),
		TicketID:   ticket.TicketID.String(),
		QRCode:     ticket.QRCode,
		IsUsed:     ticket.IsUsed,
		EventName:  ticket.Event.Title,
		TicketName: ticket.Ticket.Name,
		UsedAt:     ticket.UsedAt,
	}

	return resp, nil
}

func (s *userTicketService) ValidateTicket(qr string) (*dto.UserTicketResponse, error) {
	ticket, err := s.repo.ValidateQRCode(qr)
	if err != nil {
		return nil, response.NewBadRequest("invalid QR code")
	}
	if ticket.IsUsed {
		return nil, response.NewBadRequest("ticket already used")
	}

	resp := &dto.UserTicketResponse{
		ID:         ticket.ID.String(),
		EventID:    ticket.EventID.String(),
		TicketID:   ticket.TicketID.String(),
		QRCode:     ticket.QRCode,
		IsUsed:     ticket.IsUsed,
		EventName:  ticket.Event.Title,
		TicketName: ticket.Ticket.Name,
		UsedAt:     ticket.UsedAt,
	}

	return resp, nil
}

func (s *userTicketService) MarkTicketUsed(id string) error {
	return s.repo.MarkTicketUsed(id)
}
