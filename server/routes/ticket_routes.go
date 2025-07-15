package routes

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/middleware"

	"github.com/fiqrioemry/event_ticketing_system_app/server/handlers"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(r *gin.RouterGroup, h *handlers.TicketHandler) {
	ticket := r.Group("/tickets")
	ticket.GET("/:id", h.GetTicketByID) // TODO : Remove this endpoint later since ticket detail is not really necessary

	admin := ticket.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateTicket)
	admin.PUT("/:id", h.UpdateTicket)
	admin.DELETE("/:id", h.DeleteTicket)
}
