package routes

import (
	"server/handlers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func ReportRoutes(r *gin.RouterGroup, h *handlers.ReportHandler) {
	report := r.Group("/reports", middleware.AuthRequired(), middleware.RoleOnly("admin"))

	report.GET("/summary", h.GetSummary)
	report.GET("/orders", h.GetOrderReports)
	report.GET("/ticket-sales", h.GetTicketSalesReports)
	report.GET("/payments", h.GetPaymentReports)
	report.GET("/refunds", h.GetRefundReports)
	report.GET("/withdrawals", h.GetWithdrawalReports)
}
