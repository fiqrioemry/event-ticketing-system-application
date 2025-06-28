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

	// ========== inisialisasi Middleware ========
	r.Use(
		ginzap.Ginzap(utils.GetLogger(), time.RFC3339, true),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(100, 60*time.Second),
		middleware.LimitFileSize(12<<20),
		middleware.APIKeyGateway([]string{"/api/v1/payments/stripe/webhook"}),
	)

	// ========== inisialisasi routes ===========
	routes.InitRoutes(r, h)

	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
