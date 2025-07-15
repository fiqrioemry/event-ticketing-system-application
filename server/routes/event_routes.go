package routes

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/handlers"
	"github.com/fiqrioemry/event_ticketing_system_app/server/middleware"

	"github.com/gin-gonic/gin"
)

func EventRoutes(r *gin.RouterGroup, h *handlers.EventHandler) {
	event := r.Group("/events")

	event.GET("", h.GetAllEvents)
	event.GET("/:id", h.GetEventByID)                // TODO : Query to event detail can be optimized by separating tickets and event details
	event.GET("/:id/tickets", h.GetTicketsByEventID) // TODO : This endpoint to support optimizing event detail query, use Later after refactoring event detail query

	admin := event.Use(middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("", h.CreateEvent)
	admin.PUT("/:id", h.UpdateEventByID)
	admin.DELETE("/:id", h.DeleteEventByID)

}
