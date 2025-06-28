package routes

import (
	"server/handlers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func WithdrawalRoutes(r *gin.RouterGroup, h *handlers.WithdrawalHandler) {
	r.POST("/withdrawals", middleware.AuthRequired(), middleware.RoleOnly("user"), h.CreateWithdrawal)

	// admin endpoints
	admin := r.Group("/withdrawals", middleware.AuthRequired(), middleware.RoleOnly("admin"))
	admin.GET("", h.GetAllWithdrawals)
	admin.PATCH("/:id", h.ReviewWithdrawal)
}
