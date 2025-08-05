package handlers

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service    services.TicketService
	repository repositories.AuditLogRepository
}

func NewTicketHandler(service services.TicketService, repository repositories.AuditLogRepository) *TicketHandler {
	return &TicketHandler{service, repository}
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

	// record audit log
	auditLog := utils.BuildAuditLog(c, utils.MustGetUserID(c), "create", "ticket", newTicket)
	go h.repository.Create(c.Request.Context(), auditLog)

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

	// record audit log
	auditLog := utils.BuildAuditLog(c, utils.MustGetUserID(c), "update", "ticket", updatedTicket)
	go h.repository.Create(c.Request.Context(), auditLog)

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

	// record audit log
	auditLog := utils.BuildAuditLog(c, utils.MustGetUserID(c), "delete", "ticket", ticketID)
	go h.repository.Create(c.Request.Context(), auditLog)

	response.OK(c, "Ticket deleted successfully", ticketID)

}
