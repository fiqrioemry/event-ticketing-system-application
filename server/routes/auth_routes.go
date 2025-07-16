// internal/routes/auth_route.go
package routes

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, h *handlers.AuthHandler) {
	auth := r.Group("/auth")

	// public-endpoints
	auth.POST("/login", h.Login)
	auth.POST("/logout", h.Logout)
	auth.POST("/register", h.Register)
	auth.POST("/resend-otp", h.ResendOTP)
	auth.POST("/verify-otp", h.VerifyOTP)
	auth.POST("/refresh-token", h.RefreshToken)

}
