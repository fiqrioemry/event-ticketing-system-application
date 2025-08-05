package routes

import (
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, h *handlers.Handlers) {

	r.GET("/health", func(c *gin.Context) {
		var startTime = time.Now()
		c.JSON(200, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
			"uptime":    time.Since(startTime).Seconds(),
		})
	})

	api := r.Group("/api/v1")

	// ========= Authentication & User Management ========
	AuthRoutes(api, h.AuthHandler)
	UserRoutes(api, h.UserHandler)
	OrderRoutes(api, h.OrderHandler)
	EventRoutes(api, h.EventHandler)
	PaymentRoutes(api, h.PaymentHandler)
	TicketRoutes(api, h.TicketHandler)
	AdminRoutes(api, h.AdminHandler)
	WithdrawalRoutes(api, h.WithdrawalHandler)
	UserTicketRoutes(api, h.UserTicketHandler)

}
