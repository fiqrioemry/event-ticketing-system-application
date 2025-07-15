package main

import (
	"log"

	"github.com/fiqrioemry/event_ticketing_system_app/server/config"
	"github.com/fiqrioemry/event_ticketing_system_app/server/cron"
	"github.com/fiqrioemry/event_ticketing_system_app/server/handlers"
	"github.com/fiqrioemry/event_ticketing_system_app/server/middleware"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/routes"
	"github.com/fiqrioemry/event_ticketing_system_app/server/seeders"
	"github.com/fiqrioemry/event_ticketing_system_app/server/services"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"time"

	"github.com/fiqrioemry/go-api-toolkit/response"
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

	response.InitGin(response.InitConfig{
		Logger:              utils.GetLogger(),
		LogSuccessResponses: false,
		LogErrorResponses:   true,
	})

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

	r.SetTrustedProxies(config.AppConfig.TrustedProxies)

	// ========== inisialisasi Middleware ========
	r.Use(
		ginzap.Ginzap(utils.GetLogger(), time.RFC3339, true),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(100, 60*time.Second),
		middleware.LimitFileSize(config.AppConfig.MaxFileSize),
		middleware.APIKeyGateway(config.AppConfig.SkippedApiEndpoints),
	)

	// ========== inisialisasi routes ===========
	routes.InitRoutes(r, h)

	port := config.AppConfig.ServerPort
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
