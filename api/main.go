package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hcall/api/config"
	"hcall/api/database"
	"hcall/api/logger"
	"hcall/api/middlewares"
	"hcall/api/routes"
	"hcall/api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize config
	if err := config.LoadConfig(); err != nil {
		logger.Fatal("Main: Failed to load config", map[string]interface{}{
			"error": err.Error(),
		})
	}

	if config.AppConfig.GINMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize logger after config is loaded
	logger.InitLogger()

	// Initialize database
	database.InitDB()

	// Create Gin router
	router := gin.Default()

	// Add global middleware
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.RateLimitMiddleware(context.Background()))
	router.Use(middlewares.ValidateRequest())

	// Setup routes
	routes.SetupRoutes(router)

	// Setup all workers
	workManager := utils.GetWorkerManager()
	go workManager.StartAllWorkers()

	// Create server
	srv := &http.Server{
		Addr:    ":" + config.AppConfig.Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Main: Failed to start server", map[string]interface{}{
				"error": err.Error(),
			})
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Main: Shutting down server", nil)

	// Give outstanding requests a deadline for completion
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Main: Server forced to shutdown", map[string]interface{}{
			"error": err.Error(),
		})
	}

	logger.Info("Main: Server exiting", nil)
}
