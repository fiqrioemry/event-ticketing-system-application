package repositories

import (
	"errors"
	"server/dto"
	"server/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(data *models.Event) error
	UpdateEvent(data *models.Event) error
	DeleteEventByID(id string) error
	GetEventByID(id string) (*models.Event, error)
	WithTx(fn func(tx *gorm.DB) error) error
	GetAllEvents(params dto.EventQueryParams) ([]models.Event, int64, error)
	IsTitleTaken(title string) (bool, error)
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

func (r *eventRepository) WithTx(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *eventRepository) UpdateEvent(data *models.Event) error {
	return r.db.Save(data).Error
}

func (r *eventRepository) DeleteEventByID(id string) error {
	return r.db.Delete(&models.Event{}, "id = ?", id).Error
}

func (r *eventRepository) GetEventByID(id string) (*models.Event, error) {
	var data models.Event
	err := r.db.Preload("Tickets").First(&data, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &data, err
}

func (r *eventRepository) GetAllEvents(params dto.EventQueryParams) ([]models.Event, int64, error) {
	var events []models.Event
	var count int64

	db := r.db.Model(&models.Event{})

	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("title LIKE ?", like)
	}
	if params.Location != "" {
		db = db.Where("location = ?", params.Location)
	}

	switch params.Sort {
	case "date_asc":
		db = db.Order("date ASC")
	case "date_desc":
		db = db.Order("date DESC")
	case "created_at_asc":
		db = db.Order("created_at ASC")
	case "created_at_desc":
		db = db.Order("created_at DESC")
	default:
		db = db.Order("created_at DESC")
	}

	offset := (params.Page - 1) * params.Limit

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
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
