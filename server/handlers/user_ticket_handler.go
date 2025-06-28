package handlers

import (
	"net/http"
	"server/dto"
	"server/services"
	"server/utils"

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
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (h *UserTicketHandler) UseTicket(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.MarkTicketUsed(id); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ticket used successfully"})
}

func (h *UserTicketHandler) ValidateTicket(c *gin.Context) {
	var req dto.ValidateTicketRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}
	ticket, err := h.service.ValidateTicket(req.QRCode)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (h *UserTicketHandler) PrintTicket(c *gin.Context) {
	id := c.Param("id")
	ticket, err := h.service.GetUserTicketByID(id)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	ticketDocument := utils.GenerateTicketPDF(ticket)

	c.JSON(http.StatusOK, ticketDocument)
}
