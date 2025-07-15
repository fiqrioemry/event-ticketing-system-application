package handlers

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

type UserTicketHandler struct {
	service services.UserTicketService
}

func NewUserTicketHandler(service services.UserTicketService) *UserTicketHandler {
	return &UserTicketHandler{service}
}

func (h *UserTicketHandler) GetTicketByID(c *gin.Context) {
	id := c.Param("id")
	ticket, err := h.service.GetUserTicketByID(id)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, "Ticket retrieved successfully", ticket)
}

func (h *UserTicketHandler) UseTicket(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.MarkTicketUsed(id); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Ticket marked as used successfully", id)
}

func (h *UserTicketHandler) ValidateTicket(c *gin.Context) {
	var req dto.ValidateTicketRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	ticket, err := h.service.ValidateTicket(req.QRCode)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Ticket validated successfully", ticket)
}

func (h *UserTicketHandler) PrintTicket(c *gin.Context) {
	id := c.Param("id")
	ticket, err := h.service.GetUserTicketByID(id)
	if err != nil {
		response.Error(c, err)
		return
	}

	ticketDocument := utils.GenerateTicketPDF(ticket)

	response.OK(c, "Ticket document generated successfully", ticketDocument)
}
