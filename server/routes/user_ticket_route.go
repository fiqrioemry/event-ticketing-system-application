package routes

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/middleware"

	"github.com/fiqrioemry/event_ticketing_system_app/server/handlers"

	"github.com/gin-gonic/gin"
)

func UserTicketRoutes(r *gin.RouterGroup, h *handlers.UserTicketHandler) {
	// end-user: for getting user tickets and printing them
	user := r.Group("/user-ticket", middleware.AuthRequired(), middleware.RoleOnly("user", "admin"))
	user.GET("/:id", h.GetTicketByID)
	user.GET("/:id/print", h.PrintTicket)

	// admin/staff: for validating and marking tickets as used
	admin := r.Group("/user-ticket", middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.POST("/validate", h.ValidateTicket)
	admin.PATCH("/:id/use", h.UseTicket)
}
