package handlers

import (
	"server/dto"
	"server/services"
	"server/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service services.TicketService
}

func NewTicketHandler(service services.TicketService) *TicketHandler {
	return &TicketHandler{service}
}

func (h *TicketHandler) CreateTicket(c *gin.Context) {
	var req dto.CreateTicketRequest

	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.CreateTicket(req); err != nil {
		utils.HandleServiceError(c, err, err.Error())
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
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket updated successfully"})
}

func (h *TicketHandler) GetTicketByID(c *gin.Context) {
	ticketID := c.Param("id")

	ticket, err := h.service.GetTicketByID(ticketID)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func (h *TicketHandler) DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteTicket(id); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
