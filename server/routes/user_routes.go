// internal/routes/auth_route.go
package routes

import (
	"server/handlers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, h *handlers.UserHandler) {
	user := r.Group("/user", middleware.AuthRequired())
	user.GET("/", middleware.RoleOnly("admin"), h.GetAllUsers)
	user.GET("/:id", middleware.RoleOnly("admin"), h.GetUserDetail)
	user.GET("/me", h.GetMyProfile)
	user.PUT("/me", h.UpdateProfile)

}
