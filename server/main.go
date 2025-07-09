package main

import (
	"log"
	"os"
	"server/config"
	"server/cron"
	"server/handlers"
	"server/middleware"
	"server/repositories"
	"server/routes"
	"server/seeders"
	"server/services"
	"server/utils"

	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// TICKETING APP SERVER
// VERSION: 1.0.0
// DEPLOYMENT: docker-compose
// PORT: 5004
// DESCRIPTION: This is a server for an event ticketing system that handles user registration, event management, and payment processing.

func main() {
	// ========== Configuration =================
	config.InitConfiguration()
	utils.InitLogger()
	db := config.DB

	seeders.ResetDatabase(db)

	// ========== initialisasi layer ============
	repo := repositories.InitRepositories(db)
	s := services.InitServices(repo)
	h := handlers.InitHandlers(s)

	cronManager := cron.NewCronManager(s.PaymentService)
	cronManager.RegisterJobs()
	cronManager.Start()

	// ========== Inisialisasi gin engine =======
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		var startTime = time.Now()
		c.JSON(200, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
			"uptime":    time.Since(startTime).Seconds(),
		})
	})

	err := r.SetTrustedProxies(config.GetTrustedProxies())
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// ========== inisialisasi Middleware ========
	r.Use(
		ginzap.Ginzap(utils.GetLogger(), time.RFC3339, true),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(100, 60*time.Second),
		middleware.LimitFileSize(12<<20),
		middleware.APIKeyGateway([]string{"/api/v1/payments/stripe/webhooks"}),
	)

	// ========== inisialisasi routes ===========
	routes.InitRoutes(r, h)

	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
