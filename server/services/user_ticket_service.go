package services

import (
	"server/dto"
	customErr "server/errors"
	"server/repositories"
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
		return nil, customErr.NewNotFound("ticket not found").WithContext("ticketID", id)
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
		return nil, customErr.NewBadRequest("invalid QR code")
	}
	if ticket.IsUsed {
		return nil, customErr.NewBadRequest("ticket already used")
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
