package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.RouterGroup, h *handlers.PaymentHandler) {
	payment := r.Group("/payments")

	payment.POST("/stripe/webhooks", h.HandlePaymentNotifications)
}
