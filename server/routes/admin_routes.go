package routes

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/middleware"

	"github.com/fiqrioemry/event_ticketing_system_app/server/handlers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup, h *handlers.AdminHandler) {
	admin := r.Group("/admin", middleware.AuthRequired(), middleware.RoleOnly("admin"))

	admin.GET("/summary", h.GetSummary)
	admin.GET("/users", h.GetAllUsers)
	admin.GET("/events", h.GetAllEvents)
	admin.GET("/orders", h.GetOrderReports)
	admin.GET("/ticket-sales", h.GetTicketSalesReports)
	admin.GET("/payments", h.GetPaymentReports)
	admin.GET("/refunds", h.GetRefundReports)
	admin.GET("/withdrawals", h.GetWithdrawalReports)
}
