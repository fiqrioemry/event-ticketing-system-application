package repositories

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	DeleteEventByID(id string) error
	CreateEvent(data *models.Event) error
	UpdateEvent(data *models.Event) error
	IsTitleTaken(title string) (bool, error)
	GetEventByID(id string) (*models.Event, error)
	GetAllEvents(params dto.EventQueryParams) ([]models.Event, int64, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) CreateEvent(data *models.Event) error {
	return r.db.Create(data).Error
}

func (r *eventRepository) UpdateEvent(data *models.Event) error {
	return r.db.Save(data).Error
}

func (r *eventRepository) DeleteEventByID(id string) error {
	return r.db.Delete(&models.Event{}, "id = ?", id).Error
}

func (r *eventRepository) GetEventByID(id string) (*models.Event, error) {
	var event models.Event
	err := r.db.Preload("Tickets").First(&event, "id = ?", id).Error
	return &event, err
}

func (r *eventRepository) GetAllEvents(params dto.EventQueryParams) ([]models.Event, int64, error) {
	var events []models.Event
	var count int64

	// Mengambil semua event kecuali yang berstatus inactive
	db := r.db.Model(&models.Event{}).
		Where("events.status != ?", "inactive")

	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("events.title LIKE ? OR events.description LIKE ?", like, like)
	}

	if params.Location != "" && params.Location != "all" {
		db = db.Where("events.location = ?", params.Location)
	}

	if params.Status != "" && params.Status != "all" {
		db = db.Where("events.status = ?", params.Status)
	}

	if params.StartDate != "" {
		db = db.Where("events.date >= ?", params.StartDate)
	}

	if params.EndDate != "" {
		db = db.Where("events.date <= ?", params.EndDate)
	}

	switch params.Sort {
	case "date_asc":
		db = db.Order("events.date ASC")
	case "date_desc":
		db = db.Order("events.date DESC")
	case "title_asc":
		db = db.Order("events.title ASC")
	case "title_desc":
		db = db.Order("events.title DESC")
	default:
		db = db.Order("events.date DESC")
	}

	if params.Page <= 0 {
		params.Page = 1
	}

	if params.Limit <= 0 {
		params.Limit = 10
	}

	offset := (params.Page - 1) * params.Limit

	// Count dengan kondisi yang sama
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Query final dengan preload tickets
	if err := db.Preload("Tickets").Limit(params.Limit).Offset(offset).Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, count, nil
}
func (r *eventRepository) IsTitleTaken(title string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Event{}).Where("title = ?", title).Count(&count).Error
	return count > 0, err
}
