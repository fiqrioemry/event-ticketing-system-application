package services

import (
	"log"
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventService interface {
	DeleteEventByID(eventID string) error
	CreateEvent(req *dto.CreateEventRequest) (*models.Event, error)
	GetEventByID(id string) (*dto.EventDetailResponse, error)
	GetAllEvents(params dto.EventQueryParams) ([]dto.EventResponse, int, error)
	UpdateEvent(eventID string, req *dto.UpdateEventRequest) (*dto.EventDetailResponse, error)

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

func (s *eventService) CreateEvent(req *dto.CreateEventRequest) (*models.Event, error) { // ✅ Return event
	var createdEvent *models.Event

	err := s.repo.WithTx(func(tx *gorm.DB) error {
		today := time.Now().Truncate(24 * time.Hour)
		if !req.Date.After(today) {
			return response.NewBadRequest("Event date must be in the future")
		}

		if req.StartTime < 0 || req.StartTime > 23 {
			return response.NewBadRequest("Start time must be between 0-23")
		}

		if req.EndTime < 1 || req.EndTime > 24 {
			return response.NewBadRequest("End time must be between 1-24")
		}

		if req.StartTime >= req.EndTime {
			return response.NewBadRequest("Start time must be before end time")
		}

		duration := req.EndTime - req.StartTime
		if duration < 1 {
			return response.NewBadRequest("Event duration must be at least 1 hour")
		}

		exists, err := s.repo.IsTitleTaken(req.Title)
		if err != nil {
			return response.NewInternalServerError("Failed to check title uniqueness", err)
		}
		if exists {
			return response.NewConflict("Event title already exists")
		}

		// Create event
		event := &models.Event{
			ID:          uuid.New(),
			Image:       req.ImageURL,
			Title:       req.Title,
			Description: req.Description,
			Location:    req.Location,
			Date:        req.Date,
			StartTime:   req.StartTime,
			EndTime:     req.EndTime,
			Status:      req.Status,
		}

		if err := tx.Create(event).Error; err != nil {
			return response.NewInternalServerError("Failed to create event", err)
		}

		// Create tickets
		if len(req.Tickets) > 0 {
			var tickets []models.Ticket
			for _, ticketReq := range req.Tickets {
				tickets = append(tickets, models.Ticket{
					ID:            uuid.New(),
					EventID:       event.ID,
					Name:          ticketReq.Name,
					Price:         ticketReq.Price,
					Quota:         ticketReq.Quota,
					Limit:         ticketReq.Limit,
					Refundable:    ticketReq.Refundable,
					RefundPercent: ticketReq.RefundPercent,
				})
			}

			if err := tx.Create(&tickets).Error; err != nil {
				return response.NewInternalServerError("Failed to create tickets", err)
			}
		}

		createdEvent = event
		return nil
	})

	return createdEvent, err
}

func (s *eventService) UpdateEvent(eventID string, req *dto.UpdateEventRequest) (*dto.EventDetailResponse, error) {
	var updatedEvent *models.Event

	err := s.repo.WithTx(func(tx *gorm.DB) error {
		// Get event
		event, err := s.repo.GetEventByID(eventID)
		if err != nil {
			return response.NewInternalServerError("Failed to get event", err)
		}
		if event == nil {
			return response.NewNotFound("Event not found")
		}

		// Check if event can be updated
		if event.Status == "done" || event.Status == "cancelled" {
			return response.NewForbidden("Cannot update event with done/cancelled status")
		}

		// Date validation
		today := time.Now().Truncate(24 * time.Hour)
		if !req.Date.After(today) {
			return response.NewBadRequest("Event date must be in the future")
		}

		// Time validation
		if req.StartTime < 0 || req.StartTime > 23 {
			return response.NewBadRequest("Start time must be between 0-23")
		}
		if req.EndTime < 1 || req.EndTime > 24 {
			return response.NewBadRequest("End time must be between 1-24")
		}
		if req.StartTime >= req.EndTime {
			return response.NewBadRequest("Start time must be before end time")
		}

		// Check if tickets sold (restrict major changes)
		var totalSold int
		for _, ticket := range event.Tickets {
			totalSold += ticket.Sold
		}

		if totalSold > 0 {
			// If tickets sold, restrict date/location changes
			if !req.Date.Equal(event.Date) || req.Location != event.Location {
				return response.NewBadRequest("Cannot update date/location after tickets are sold")
			}
		}

		// Check title uniqueness (if changed)
		if req.Title != event.Title {
			exists, err := s.repo.IsTitleTaken(req.Title)
			if err != nil {
				return response.NewInternalServerError("Failed to check title uniqueness", err)
			}
			if exists {
				return response.NewConflict("Event title already exists")
			}
		}

		// Handle image update
		if req.ImageURL != "" {
			if err := utils.DeleteFromCloudinary(event.Image); err != nil {
				log.Printf("Failed to delete old image: %v", err)
			}
			event.Image = req.ImageURL
		}

		// Update event fields
		event.Title = req.Title
		event.Description = req.Description
		event.Location = req.Location
		event.Date = req.Date
		event.StartTime = req.StartTime
		event.EndTime = req.EndTime
		event.Status = req.Status

		// Save updated event
		if err := tx.Save(event).Error; err != nil {
			return response.NewInternalServerError("Failed to update event", err)
		}

		// Get fresh event with updated data and tickets
		freshEvent, err := s.repo.GetEventByIDWithTx(tx, eventID)
		if err != nil {
			return response.NewInternalServerError("Failed to get updated event", err)
		}

		updatedEvent = freshEvent
		return nil
	})

	if err != nil {
		return nil, err
	}

	var tickets []dto.TicketResponse
	for _, ticket := range updatedEvent.Tickets {
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

	eventResponse := &dto.EventDetailResponse{
		ID:          updatedEvent.ID.String(),
		Title:       updatedEvent.Title,
		Image:       updatedEvent.Image,
		Description: updatedEvent.Description,
		Location:    updatedEvent.Location,
		Date:        updatedEvent.Date,
		StartTime:   updatedEvent.StartTime,
		EndTime:     updatedEvent.EndTime,
		Status:      updatedEvent.Status,
		Tickets:     tickets,
		CreatedAt:   updatedEvent.CreatedAt,
	}

	return eventResponse, nil
}

func (s *eventService) GetAllEvents(params dto.EventQueryParams) ([]dto.EventResponse, int, error) {

	list, total, err := s.repo.GetAllEvents(params)
	if err != nil {
		return nil, 0, response.NewInternalServerError("Failed to get events", err)

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
			Status:      item.Status,
			Date:        item.Date,
			CreatedAt:   item.CreatedAt,
		})
	}

	return result, int(total), nil
}

func (s *eventService) GetEventByID(id string) (*dto.EventDetailResponse, error) {
	event, err := s.repo.GetEventByID(id)
	if event == nil || err != nil {
		return nil, response.NewNotFound("event not found")
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
		Date:        event.Date,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Status:      event.Status,
		Tickets:     tickets,
		CreatedAt:   event.CreatedAt,
	}, nil
}

func (s *eventService) GetAllTicketsByEventID(eventID string) ([]dto.TicketResponse, error) {
	tickets, err := s.ticket.GetAllTicketsByEventID(eventID)
	if err != nil {
		return nil, response.NewNotFound("event not found")
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

func (s *eventService) DeleteEventByID(eventID string) error {
	event, err := s.repo.GetEventByID(eventID)
	if event == nil || err != nil {
		return response.NewNotFound("event not found")
	}

	if event.Status == "done" || event.Status == "ongoing" {
		return response.NewForbidden("cannot delete event with done/ongoing status")
	}

	for _, ticket := range event.Tickets {
		if ticket.Sold > 0 {
			return response.NewForbidden("cannot delete event with sold tickets")
		}
	}

	if event.Image != "" {
		err = utils.DeleteFromCloudinary(event.Image)
		if err != nil {
			return response.NewInternalServerError("failed to delete image", err)
		}
	}

	return s.repo.DeleteEventByID(eventID)
}
