package handlers

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/fiqrioemry/go-api-toolkit/response"
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
		response.Error(c, response.NewBadRequest("Event ID is required"))
		return
	}

	var req dto.CreateTicketRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	newTicket, err := h.service.CreateTicket(req, eventID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, "Ticket created successfully", newTicket)
}

func (h *TicketHandler) UpdateTicket(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateTicketRequest

	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	updatedTicket, err := h.service.UpdateTicket(id, req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Ticket updated successfully", updatedTicket)
}

func (h *TicketHandler) GetTicketByID(c *gin.Context) {
	ticketID := c.Param("id")

	ticket, err := h.service.GetTicketByID(ticketID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Ticket retrieved successfully", ticket)
}

func (h *TicketHandler) DeleteTicket(c *gin.Context) {
	ticketID := c.Param("id")

	if err := h.service.DeleteTicket(ticketID); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Ticket deleted successfully", ticketID)

}
