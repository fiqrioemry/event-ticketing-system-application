package handlers

import (
	"net/http"
	"server/dto"
	"server/services"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service services.EventService
}

func NewEventHandler(service services.EventService) *EventHandler {
	return &EventHandler{service}
}

func (h *EventHandler) GetAllEvents(c *gin.Context) {
	var params dto.EventQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}
	data, pagination, err := h.service.GetAllEvents(params)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data, "pagination": pagination})
}

func (h *EventHandler) GetEventByID(c *gin.Context) {
	id := c.Param("id")
	data, err := h.service.GetEventByID(id)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var req dto.CreateEventRequest
	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	imageURL, err := utils.UploadImageWithValidation(req.Image)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	req.ImageURL = imageURL
	if err := h.service.CreateEvent(req); err != nil {
		utils.CleanupImageOnError(imageURL)
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event Created"})
}

func (h *EventHandler) GetTicketsByEventID(c *gin.Context) {
	eventID := c.Param("id")

	tickets, err := h.service.GetAllTicketsByEventID(eventID)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func (h *EventHandler) UpdateEventByID(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateEventRequest
	if !utils.BindAndValidateForm(c, &req) {
		return
	}
	if req.Image != nil && req.Image.Filename != "" {
		imageURL, err := utils.UploadImageWithValidation(req.Image)
		if err != nil {
			utils.HandleServiceError(c, err, err.Error())
			return
		}
		req.ImageURL = imageURL
	}

	if err := h.service.UpdateEvent(id, req); err != nil {
		utils.CleanupImageOnError(req.ImageURL)
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func (h *EventHandler) DeleteEventByID(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteEventByID(id); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
