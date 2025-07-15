package handlers

import (
	"net/http"

	"github.com/fiqrioemry/event_ticketing_system_app/server/config"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/fiqrioemry/go-api-toolkit/response"
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
		response.Error(c, err)
		return
	}

	sigHeader := c.GetHeader("Stripe-Signature")

	event, err := webhook.ConstructEventWithOptions(body, sigHeader, config.AppConfig.StripeWebhookSecret, webhook.ConstructEventOptions{
		IgnoreAPIVersionMismatch: true,
	})
	if err != nil {
		response.Error(c, err)
		return
	}

	if err := h.service.StripeWebhookNotification(event); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Payment notification processed successfully", nil)
}

// TODO : // 1. Add a method to create a payment, so admin can do it manually based on requests
// TODO : // 2. Add a method to update payment status, so admin can do it manually based on requests
