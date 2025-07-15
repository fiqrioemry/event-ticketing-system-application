package handlers

import (
	"encoding/json"

	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/fiqrioemry/go-api-toolkit/pagination"
	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type EventHandler struct {
	service services.EventService
}

func NewEventHandler(service services.EventService) *EventHandler {
	return &EventHandler{service}
}

func (h *EventHandler) GetAllEvents(c *gin.Context) {
	var params dto.EventQueryParams
	// bind query params
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	// apply pagination defaults
	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	// fetch events data
	data, total, err := h.service.GetAllEvents(params)
	if err != nil {
		response.Error(c, err)
		return
	}

	// build pagination meta
	pag := pagination.Build(params.Page, params.Limit, total)

	response.OKWithPagination(c, "events retrieved successfully", data, pag)
}

func (h *EventHandler) GetEventByID(c *gin.Context) {
	// extract event ID
	id := c.Param("id")
	// fetch event data
	data, err := h.service.GetEventByID(id)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, "Event retrieved successfully", data)
}

func buildValidationError(err error) *response.AppError {
	validationErr := response.NewBadRequest("Validation failed")
	validationErr.WithContext("details", err.Error())
	return validationErr
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	// extract form image
	image, err := c.FormFile("image")
	if err != nil {
		response.Error(c, response.BadRequest("Event image is required"))
		return
	}

	// extract form data
	dataStr := c.PostForm("data")
	if dataStr == "" {
		response.Error(c, response.BadRequest("Event data is required"))
		return
	}

	// parse JSON data
	var req dto.CreateEventRequest
	if err := json.Unmarshal([]byte(dataStr), &req); err != nil {
		response.Error(c, response.BadRequest("Invalid JSON data format: "+err.Error()))
		return
	}

	req.Image = image

	// validate request struct
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			validationErr := buildValidationError(validationErrors)
			response.Error(c, validationErr)
			return
		}
		response.Error(c, response.NewBadRequest("Validation failed: "+err.Error()))
		return
	}

	// upload image file
	imageURL, err := utils.UploadImageWithValidation(req.Image)
	if err != nil {
		response.Error(c, err)
		return
	}
	req.ImageURL = imageURL

	// create event record
	createdEvent, err := h.service.CreateEvent(&req)
	if err != nil {
		// cleanup failed upload
		utils.CleanupImageOnError(imageURL)
		response.Error(c, err)
		return
	}

	response.Created(c, "Event and tickets created successfully", createdEvent)
}

func (h *EventHandler) UpdateEventByID(c *gin.Context) {
	// extract event ID
	id := c.Param("id")

	// extract form data
	dataStr := c.PostForm("data")
	if dataStr == "" {
		response.Error(c, response.NewBadRequest("Event data is required"))
		return
	}

	// parse JSON data
	var req dto.UpdateEventRequest
	if err := json.Unmarshal([]byte(dataStr), &req); err != nil {
		response.Error(c, response.NewBadRequest("Invalid JSON data format: "+err.Error()))
		return
	}

	// check optional image
	image, err := c.FormFile("image")
	if err == nil && image != nil && image.Filename != "" {
		req.Image = image
	}

	// validate request struct
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			validationErr := buildValidationError(validationErrors)
			response.Error(c, validationErr)
			return
		}
		response.Error(c, response.NewBadRequest("Validation failed: "+err.Error()))
		return
	}

	// upload new image
	if req.Image != nil && req.Image.Filename != "" {
		imageURL, err := utils.UploadImageWithValidation(req.Image)
		if err != nil {
			response.Error(c, err)
			return
		}
		req.ImageURL = imageURL
	}

	// update event record
	updatedEvent, err := h.service.UpdateEvent(id, &req)
	if err != nil {
		// cleanup failed upload
		if req.ImageURL != "" {
			utils.CleanupImageOnError(req.ImageURL)
		}
		response.Error(c, err)
		return
	}

	response.OK(c, "Event updated successfully", updatedEvent)
}

func (h *EventHandler) GetTicketsByEventID(c *gin.Context) {
	// extract event ID
	eventID := c.Param("id")

	// fetch tickets data
	tickets, err := h.service.GetAllTicketsByEventID(eventID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, "Tickets retrieved successfully", tickets)
}

func (h *EventHandler) DeleteEventByID(c *gin.Context) {
	// extract event ID
	eventID := c.Param("id")
	// delete event record
	if err := h.service.DeleteEventByID(eventID); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Event deleted successfully", eventID)
}
