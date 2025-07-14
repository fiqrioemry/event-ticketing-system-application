package handlers

import (
	"encoding/json"
	"net/http"
	"server/dto"
	"server/errors"
	customErr "server/errors"
	"server/services"
	"server/utils"

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
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	data, total, err := h.service.GetAllEvents(params)
	if err != nil {
		response.Error(c, err)
		return
	}

	pag := pagination.Build(params.Page, params.Limit, total)

	response.OKWithPagination(c, "events retrieved successfully", data, pag)
}

func (h *EventHandler) GetEventByID(c *gin.Context) {
	id := c.Param("id")
	data, err := h.service.GetEventByID(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func buildValidationError(err error) *errors.AppError {
	validationErr := errors.NewBadRequest("Validation failed")
	validationErr.WithContext("details", err.Error())
	return validationErr
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	// 1. Parse image from form
	image, err := c.FormFile("image")
	if err != nil {
		utils.HandleError(c, customErr.NewBadRequest("Image is required"))
		return
	}

	// 2. Parse JSON data from form
	dataStr := c.PostForm("data")
	if dataStr == "" {
		utils.HandleError(c, customErr.NewBadRequest("Event data is required"))
		return
	}

	// 3. Unmarshal JSON data
	var req dto.CreateEventRequest
	if err := json.Unmarshal([]byte(dataStr), &req); err != nil {
		utils.HandleError(c, customErr.NewBadRequest("Invalid JSON data format: "+err.Error()))
		return
	}

	// 4. Attach image to request (after JSON parsing)
	req.Image = image

	// 5. Manual validation using validator
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			validationErr := buildValidationError(validationErrors)
			utils.HandleError(c, validationErr)
			return
		}
		utils.HandleError(c, customErr.NewBadRequest("Validation failed: "+err.Error()))
		return
	}

	// 6. Upload image
	imageURL, err := utils.UploadImageWithValidation(req.Image)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	req.ImageURL = imageURL

	// 7. Create event + tickets
	event, err := h.service.CreateEvent(&req)
	if err != nil {
		utils.CleanupImageOnError(imageURL)
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Event and tickets created successfully",
		"data": gin.H{
			"event_id":      event.ID,
			"title":         event.Title,
			"status":        event.Status,
			"tickets_count": len(req.Tickets),
		},
	})
}
func (h *EventHandler) UpdateEventByID(c *gin.Context) {
	id := c.Param("id")

	// 1. Parse JSON data from form (required)
	dataStr := c.PostForm("data")
	if dataStr == "" {
		utils.HandleError(c, customErr.NewBadRequest("Event data is required"))
		return
	}

	// 2. Unmarshal JSON data
	var req dto.UpdateEventRequest
	if err := json.Unmarshal([]byte(dataStr), &req); err != nil {
		utils.HandleError(c, customErr.NewBadRequest("Invalid JSON data format: "+err.Error()))
		return
	}

	// 3. Parse image from form (optional)
	image, err := c.FormFile("image")
	if err == nil && image != nil && image.Filename != "" {
		// Image provided, attach to request
		req.Image = image
	}
	// If err != nil, it means no image provided (which is OK for update)

	// 4. Manual validation using validator
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			validationErr := buildValidationError(validationErrors)
			utils.HandleError(c, validationErr)
			return
		}
		utils.HandleError(c, customErr.NewBadRequest("Validation failed: "+err.Error()))
		return
	}

	// 5. Handle image upload (if provided)
	if req.Image != nil && req.Image.Filename != "" {
		imageURL, err := utils.UploadImageWithValidation(req.Image)
		if err != nil {
			utils.HandleError(c, err)
			return
		}
		req.ImageURL = imageURL
	}

	// 6. Update event
	if err := h.service.UpdateEvent(id, &req); err != nil {
		if req.ImageURL != "" {
			utils.CleanupImageOnError(req.ImageURL)
		}
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Event updated successfully",
	})
}

func (h *EventHandler) GetTicketsByEventID(c *gin.Context) {
	eventID := c.Param("id")

	tickets, err := h.service.GetAllTicketsByEventID(eventID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func (h *EventHandler) DeleteEventByID(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteEventByID(id); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
