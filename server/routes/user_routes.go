// internal/routes/auth_route.go
package routes

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/handlers"
	"github.com/fiqrioemry/event_ticketing_system_app/server/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, h *handlers.UserHandler) {

	user := r.Group("/user", middleware.AuthRequired())
	user.GET("/me", h.GetMyProfile)
	user.PUT("/me", h.UpdateProfile)
	user.PUT("/change-password", h.ChangePassword)
}
