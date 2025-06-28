package routes

import (
	"server/handlers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func ReportRoutes(r *gin.RouterGroup, h *handlers.ReportHandler) {
	report := r.Group("/reports", middleware.AuthRequired(), middleware.RoleOnly("admin"))
	report.GET("/transaction", h.GetAllTransactions)
	report.GET("/transaction/:id", h.GetTransactionByID)
}
