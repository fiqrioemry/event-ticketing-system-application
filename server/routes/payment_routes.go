package routes

import (
	"server/handlers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.RouterGroup, h *handlers.PaymentHandler) {
	payment := r.Group("/payments", middleware.AuthRequired())

	payment.POST("/stripe/webhook", h.HandlePaymentNotifications)
}
