package handlers

import (
	"net/http"
	"os"
	"server/services"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75/webhook"
)

type PaymentHandler struct {
	service services.PaymentService
}

func NewPaymentHandler(service services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service}
}

func (h *PaymentHandler) HandlePaymentNotifications(c *gin.Context) {
	const MaxBodyBytes = int64(65536)

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	body, err := c.GetRawData()
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	sigHeader := c.GetHeader("Stripe-Signature")

	event, err := webhook.ConstructEventWithOptions(body, sigHeader, os.Getenv("STRIPE_WEBHOOK_SECRET"), webhook.ConstructEventOptions{
		IgnoreAPIVersionMismatch: true,
	})
	if err != nil {
		utils.HandleServiceError(c, err, "Invalid Signatures")
		return
	}

	if err := h.service.StripeWebhookNotification(event); err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment received successfully"})
}

// TODO : // 1. Add a method to create a payment, so admin can do it manually based on requests
// TODO : // 2. Add a method to update payment status, so admin can do it manually based on requests
