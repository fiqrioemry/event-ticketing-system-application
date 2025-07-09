package services

import (
	"log"
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
	CreateEvent(req *dto.CreateEventRequest) (*models.Event, error)
	GetEventByID(id string) (*dto.EventDetailResponse, error)
	GetAllEvents(params dto.EventQueryParams) ([]dto.EventResponse, *dto.PaginationResponse, error)
	UpdateEvent(eventID string, req *dto.UpdateEventRequest) error

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
		// Date validation
		today := time.Now().Truncate(24 * time.Hour)
		if !req.Date.After(today) {
			return customErr.NewBadRequest("Event date must be in the future")
		}

		// Time validation - Fix logic
		if req.StartTime < 0 || req.StartTime > 23 {
			return customErr.NewBadRequest("Start time must be between 0-23")
		}

		if req.EndTime < 1 || req.EndTime > 24 {
			return customErr.NewBadRequest("End time must be between 1-24")
		}

		if req.StartTime >= req.EndTime {
			return customErr.NewBadRequest("Start time must be before end time")
		}

		// ✅ Better duration validation
		duration := req.EndTime - req.StartTime
		if duration < 1 {
			return customErr.NewBadRequest("Event duration must be at least 1 hour")
		}

		// ✅ Check title uniqueness
		exists, err := s.repo.IsTitleTaken(req.Title)
		if err != nil {
			return customErr.NewInternalServerError("Failed to check title uniqueness", err)
		}
		if exists {
			return customErr.NewConflict("Event title already exists")
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
			return customErr.NewInternalServerError("Failed to create event", err)
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
				return customErr.NewInternalServerError("Failed to create tickets", err)
			}
		}

		createdEvent = event
		return nil
	})

	return createdEvent, err
}

func (s *eventService) UpdateEvent(eventID string, req *dto.UpdateEventRequest) error {
	return s.repo.WithTx(func(tx *gorm.DB) error {
		// Get event
		event, err := s.repo.GetEventByID(eventID)
		if err != nil {
			return customErr.NewInternalServerError("Failed to get event", err)
		}
		if event == nil {
			return customErr.NewNotFound("Event not found")
		}

		// Check if event can be updated
		if event.Status == "done" || event.Status == "cancelled" {
			return customErr.NewForbidden("Cannot update event with done/cancelled status")
		}

		// Date validation
		today := time.Now().Truncate(24 * time.Hour)
		if !req.Date.After(today) {
			return customErr.NewBadRequest("Event date must be in the future")
		}

		// Time validation
		if req.StartTime < 0 || req.StartTime > 23 {
			return customErr.NewBadRequest("Start time must be between 0-23")
		}
		if req.EndTime < 1 || req.EndTime > 24 {
			return customErr.NewBadRequest("End time must be between 1-24")
		}
		if req.StartTime >= req.EndTime {
			return customErr.NewBadRequest("Start time must be before end time")
		}

		// Check if tickets sold (restrict major changes)
		var totalSold int
		for _, ticket := range event.Tickets {
			totalSold += ticket.Sold
		}

		if totalSold > 0 {
			// If tickets sold, restrict date/location changes
			if !req.Date.Equal(event.Date) || req.Location != event.Location {
				return customErr.NewBadRequest("Cannot update date/location after tickets are sold")
			}
		}

		// Check title uniqueness (if changed)
		if req.Title != event.Title {
			exists, err := s.repo.IsTitleTaken(req.Title)
			if err != nil {
				return customErr.NewInternalServerError("Failed to check title uniqueness", err)
			}
			if exists {
				return customErr.NewConflict("Event title already exists")
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

		return tx.Save(event).Error // ✅ Use tx.Save instead of repo method
	})
}
func (s *eventService) GetAllEvents(params dto.EventQueryParams) ([]dto.EventResponse, *dto.PaginationResponse, error) {

	list, total, err := s.repo.GetAllEvents(params)
	if err != nil {
		return nil, nil, customErr.NewInternalServerError("failed to retrieve events", err)

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

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return result, pagination, nil
}

func (s *eventService) GetEventByID(id string) (*dto.EventDetailResponse, error) {
	event, err := s.repo.GetEventByID(id)
	if event == nil || err != nil {
		return nil, customErr.NewNotFound("event not found")
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
		return nil, customErr.NewNotFound("event not found")
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
			return customErr.NewInternalServerError("failed to delete image", err)
		}
	}

	return s.repo.DeleteEventByID(eventID)
}
