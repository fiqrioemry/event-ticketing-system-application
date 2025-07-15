package config

import (
	"log"

	"github.com/stripe/stripe-go/v75"
)

func InitStripe() {
	stripe.Key = AppConfig.StripeSecretKey

	log.Println("Stripe client initialized")
}
