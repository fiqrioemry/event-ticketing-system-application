package cron

import (
	"log"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/robfig/cron/v3"
)

type CronManager struct {
	c              *cron.Cron
	paymentService services.PaymentService
}

func NewCronManager(
	payment services.PaymentService,
) *CronManager {
	return &CronManager{
		c:              cron.New(cron.WithSeconds()),
		paymentService: payment,
	}
}

func (cm *CronManager) RegisterJobs() {
	// Update payment pending → failed (check for every 15 ninutes) )
	cm.c.AddFunc("0 */15 * * * *", func() {
		log.Println("Cron: Checking expired pending payments...")
		if err := cm.paymentService.ExpireOldPendingPayments(); err != nil {
			log.Println("Error expiring payments:", err)
		} else {
			log.Println("Payment status updated (pending → failed)")
		}
	})
}
func (cm *CronManager) Start() {
	cm.c.Start()
	log.Println("Cron Manager started")
}
