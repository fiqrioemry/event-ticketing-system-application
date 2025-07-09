package handlers

import (
	"server/dto"
	"server/services"
	"server/utils"

	"net/http"
	customErr "server/errors"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service services.TicketService
}

func NewTicketHandler(service services.TicketService) *TicketHandler {
	return &TicketHandler{service}
}

func (h *TicketHandler) CreateTicket(c *gin.Context) {
	eventID := c.Param("id")
	if eventID == "" {
		utils.HandleError(c, customErr.NewBadRequest("Event ID is required"))
		return
	}

	var req dto.CreateTicketRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.CreateTicket(req, eventID); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Ticket created successfully"})
}

func (h *TicketHandler) UpdateTicket(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateTicketRequest

	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.UpdateTicket(id, req); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket updated successfully"})
}

func (h *TicketHandler) GetTicketByID(c *gin.Context) {
	ticketID := c.Param("id")

	ticket, err := h.service.GetTicketByID(ticketID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func (h *TicketHandler) DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteTicket(id); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
