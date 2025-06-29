package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, h *handlers.Handlers) {
	api := r.Group("/api/v1")

	// ========= Authentication & User Management ========
	AuthRoutes(api, h.AuthHandler)
	UserRoutes(api, h.UserHandler)
	OrderRoutes(api, h.OrderHandler)
	EventRoutes(api, h.EventHandler)
	PaymentRoutes(api, h.PaymentHandler)
	TicketRoutes(api, h.TicketHandler)
	ReportRoutes(api, h.ReportHandler)
	WithdrawalRoutes(api, h.WithdrawalHandler)
	UserTicketRoutes(api, h.UserTicketHandler)

}
