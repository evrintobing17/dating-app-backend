// main.go
package main

import (
	"log"

	"github.com/evrintobing17/dating-app-go/config"
	"github.com/evrintobing17/dating-app-go/internal/repository"
	"github.com/gin-gonic/gin"

	"github.com/evrintobing17/dating-app-go/internal/middleware"

	authDelivery "github.com/evrintobing17/dating-app-go/internal/module/auth/delivery/http"
	authRepository "github.com/evrintobing17/dating-app-go/internal/module/auth/repository"
	authUsecasse "github.com/evrintobing17/dating-app-go/internal/module/auth/usecase"

	premiumDelivery "github.com/evrintobing17/dating-app-go/internal/module/premium/delivery/http"
	premiumRepository "github.com/evrintobing17/dating-app-go/internal/module/premium/repository"
	premiumUsecasse "github.com/evrintobing17/dating-app-go/internal/module/premium/usecase"

	swipeDelivery "github.com/evrintobing17/dating-app-go/internal/module/swipe/delivery/http"
	swipeRepository "github.com/evrintobing17/dating-app-go/internal/module/swipe/repository"
	swipeUsecasse "github.com/evrintobing17/dating-app-go/internal/module/swipe/usecase"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database connection
	db, err := repository.NewDatabase(cfg.DB)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Run database migrations
	if err := repository.RunMigrations(cfg.DB); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	r := gin.New()

	// Initialize Redis connection
	redisClient := repository.NewRedisClient(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)
	defer redisClient.Close()

	// Initialize Repo
	authRepo := authRepository.NewAuthRepository(db)
	premiumRepo := premiumRepository.NewPremiumRepository(db)
	swipeRepo := swipeRepository.NewSwipeRepository(db)

	// Initialize Usecase
	authUC := authUsecasse.NewAuthUsecase(authRepo)
	premiumUC := premiumUsecasse.NewPremiumUsecase(premiumRepo, authRepo)
	swipeUC := swipeUsecasse.NewSwipeUsecase(swipeRepo, redisClient)

	// Initialize Middleware
	middleware := middleware.NewAuthMiddleware(authRepo)

	// Initialize Handler
	authDelivery.NewAuthHandler(r, authUC)
	premiumDelivery.NewPremiumHandler(r, premiumUC, middleware)
	swipeDelivery.NewSwipeHandler(r, swipeUC, middleware)

	// Initialize API handlers
	// api.SetupRoutes(r, swipeService, premiumService)

	// Start server
	log.Printf("Server running on port %s", cfg.Server.Port)
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
