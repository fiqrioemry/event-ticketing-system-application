package services

import (
	"server/dto"
	customErr "server/errors"
	"server/models"
	"server/repositories"
	"server/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventService interface {
	DeleteEventByID(eventID string) error
	CreateEvent(req dto.CreateEventRequest) error
	GetEventByID(id string) (*dto.EventDetailResponse, error)
	GetAllEvents(params dto.EventQueryParams) ([]dto.EventResponse, *dto.PaginationResponse, error)
	UpdateEvent(eventID string, req dto.UpdateEventRequest) error

	// GET
	GetAllTicketsByEventID(eventID string) ([]dto.TicketResponse, error)
}

type eventService struct {
	repo   repositories.EventRepository
	ticket repositories.TicketRepository
}

func NewEventService(repo repositories.EventRepository, ticket repositories.TicketRepository) EventService {
	return &eventService{repo, ticket}
}

func (s *eventService) CreateEvent(req dto.CreateEventRequest) error {
	return s.repo.WithTx(func(tx *gorm.DB) error {
		parsedDate, err := utils.ParseDate(req.Date)
		if err != nil {
			return err
		}

		today := time.Now().Truncate(24 * time.Hour)
		if !parsedDate.After(today) {
			return customErr.NewBadRequest("event date must be in the future")
		}

		event := &models.Event{
			ID:          uuid.New(),
			Image:       req.ImageURL,
			Title:       req.Title,
			Description: req.Description,
			Location:    req.Location,
			Date:        parsedDate,
			StartTime:   req.StartTime,
			EndTime:     req.EndTime,
			Status:      req.Status,
		}

		if err := tx.Create(event).Error; err != nil {
			return customErr.NewInternal("failed to create event", err)
		}

		var tickets []models.Ticket
		for _, cat := range req.Tickets {
			if cat.Quota < 0 || cat.Price < 0 {
				return customErr.NewBadRequest("invalid price or quota")
			}

			tickets = append(tickets, models.Ticket{
				ID:         uuid.New(),
				EventID:    event.ID,
				Name:       cat.Name,
				Price:      cat.Price,
				Quota:      cat.Quota,
				Limit:      cat.Limit,
				Refundable: cat.Refundable,
			})
		}

		if err := tx.Create(&tickets).Error; err != nil {
			return customErr.NewInternal("failed to create tickets", err)
		}

		return nil
	})
}

func (s *eventService) GetAllEvents(params dto.EventQueryParams) ([]dto.EventResponse, *dto.PaginationResponse, error) {
	list, total, err := s.repo.GetAllEvents(params)
	if err != nil {
		return nil, nil, err
	}
	var result []dto.EventResponse
	for _, item := range list {
		totalQuota := 0
		startPrice := 0.0
		for _, ticket := range item.Tickets {
			totalQuota += ticket.Quota
			if startPrice == 0.0 || float64(ticket.Price) < startPrice {
				startPrice = float64(ticket.Price)
			}
		}

		isAvailable := (item.Status == "active" || item.Status == "ongoing") && totalQuota > 0

		result = append(result, dto.EventResponse{
			ID:          item.ID.String(),
			Image:       item.Image,
			Title:       item.Title,
			Description: item.Description,
			Location:    item.Location,
			StartPrice:  startPrice,
			IsAvailable: isAvailable,
			StartTime:   item.StartTime,
			EndTime:     item.EndTime,
			Date:        item.Date.Format("2006-01-02"),
			CreatedAt:   item.CreatedAt.Format("2006-01-02"),
		})
	}
	pagination := utils.Paginate(total, params.Page, params.Limit)
	return result, pagination, nil
}

func (s *eventService) GetEventByID(id string) (*dto.EventDetailResponse, error) {
	event, err := s.repo.GetEventByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, customErr.NewNotFound("event not found")
		}
		return nil, customErr.NewInternal("failed to get event", err)
	}

	var tickets []dto.TicketResponse
	for _, ticket := range event.Tickets {
		tickets = append(tickets, dto.TicketResponse{
			ID:         ticket.ID.String(),
			EventID:    ticket.EventID.String(),
			Name:       ticket.Name,
			Price:      ticket.Price,
			Quota:      ticket.Quota,
			Limit:      ticket.Limit,
			Sold:       ticket.Sold,
			Refundable: ticket.Refundable,
		})
	}
	return &dto.EventDetailResponse{
		ID:          event.ID.String(),
		Title:       event.Title,
		Image:       event.Image,
		Description: event.Description,
		Location:    event.Location,
		Date:        event.Date.Format("2006-01-02"),
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Status:      event.Status,
		Tickets:     tickets,
		CreatedAt:   event.CreatedAt.Format("2006-01-02"),
	}, nil
}

func (s *eventService) GetAllTicketsByEventID(eventID string) ([]dto.TicketResponse, error) {
	tickets, err := s.ticket.GetAllTicketsByEventID(eventID)
	if err != nil {
		return nil, customErr.NewNotFound("ticket not found")
	}

	var responses []dto.TicketResponse
	for _, ticket := range tickets {
		responses = append(responses, dto.TicketResponse{
			ID:         ticket.ID.String(),
			EventID:    ticket.EventID.String(),
			Name:       ticket.Name,
			Price:      ticket.Price,
			Limit:      ticket.Limit,
			Quota:      ticket.Quota,
			Sold:       ticket.Sold,
			Refundable: ticket.Refundable,
		})
	}

	return responses, nil

}

func (s *eventService) UpdateEvent(eventID string, req dto.UpdateEventRequest) error {
	return s.repo.WithTx(func(tx *gorm.DB) error {
		event, err := s.repo.GetEventByID(eventID)
		if err != nil {
			return customErr.NewNotFound("event not found")
		}

		if event.Status == "done" || event.Status == "cancelled" {
			return customErr.NewForbidden("cannot update event with done/cancelled status")
		}

		parsedDate, err := utils.ParseDate(req.Date)
		if err != nil {
			return err
		}

		today := time.Now().Truncate(24 * time.Hour)
		if !parsedDate.After(today) {
			return customErr.NewBadRequest("event date must be in the future")
		}

		for _, ticket := range event.Tickets {
			if ticket.Sold > 0 && (!parsedDate.Equal(event.Date) || req.Location != event.Location) {
				return customErr.NewBadRequest("cannot update date/location after tickets sold")
			}

		}

		if req.Title != event.Title {
			taken, err := s.repo.IsTitleTaken(req.Title)
			if err != nil {
				return customErr.NewInternal("title check failed", err)
			}
			if taken {
				return customErr.NewBadRequest("event title already exists")
			}
		}

		if req.ImageURL != "" {
			if err := utils.DeleteFromCloudinary(event.Image); err != nil {
				return customErr.NewInternal("failed to delete old image", err)
			}
			event.Image = req.ImageURL
		}

		event.Title = req.Title
		event.Date = parsedDate
		event.Status = req.Status
		event.Location = req.Location
		event.EndTime = req.EndTime
		event.StartTime = req.StartTime
		event.Description = req.Description

		return s.repo.UpdateEvent(event)
	})
}

func (s *eventService) DeleteEventByID(eventID string) error {
	event, err := s.repo.GetEventByID(eventID)
	if err != nil {
		return customErr.NewNotFound("event not found")
	}

	if event.Status == "done" || event.Status == "ongoing" {
		return customErr.NewForbidden("cannot delete event with done/ongoing status")
	}

	for _, ticket := range event.Tickets {
		if ticket.Sold > 0 {
			return customErr.NewForbidden("cannot delete event with sold tickets")
		}
	}

	if event.Image != "" {
		err = utils.DeleteFromCloudinary(event.Image)
		if err != nil {
			return customErr.NewInternal("failed to delete image", err)
		}
	}

	return s.repo.DeleteEventByID(eventID)
}
