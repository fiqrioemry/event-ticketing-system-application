package handlers

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/fiqrioemry/go-api-toolkit/pagination"
	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service    services.EventService
	repository repositories.AuditLogRepository
}

func NewEventHandler(service services.EventService, repository repositories.AuditLogRepository) *EventHandler {
	return &EventHandler{service, repository}
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

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var req dto.CreateEventRequest

	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	imageURL, err := utils.UploadImageWithValidation(req.Image)
	if err != nil {
		response.Error(c, err)
		return
	}

	req.ImageURL = imageURL

	// create event record
	createdEvent, err := h.service.CreateEvent(&req)
	if err != nil {
		utils.CleanupImageOnError(imageURL)
		response.Error(c, err)
		return
	}

	// record audit log
	auditLog := utils.BuildAuditLog(c, utils.MustGetUserID(c), "create", "event", req)

	go h.repository.Create(c.Request.Context(), auditLog)

	response.Created(c, "Event and tickets created successfully", createdEvent)
}

func (h *EventHandler) UpdateEventByID(c *gin.Context) {
	eventID := c.Param("id")
	var req dto.UpdateEventRequest

	if !utils.BindAndValidateForm(c, &req) {
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
	updatedEvent, err := h.service.UpdateEvent(eventID, &req)
	if err != nil {
		if req.ImageURL != "" {
			utils.CleanupImageOnError(req.ImageURL)
		}
		response.Error(c, err)
		return
	}

	// record audit log
	auditLog := utils.BuildAuditLog(c, utils.MustGetUserID(c), "update", "event", req)

	go h.repository.Create(c.Request.Context(), auditLog)

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
	eventID := c.Param("id")

	// delete event record
	if err := h.service.DeleteEventByID(eventID); err != nil {
		response.Error(c, err)
		return
	}

	// record audit log
	auditLog := utils.BuildAuditLog(c, utils.MustGetUserID(c), "delete", "event", eventID)

	go h.repository.Create(c.Request.Context(), auditLog)

	response.OK(c, "Event deleted successfully", eventID)
}
