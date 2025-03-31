package main

import (
	"log"

	"hcall/api/config"
	"hcall/api/database"
	"hcall/api/middlewares"
	"hcall/api/routes"
	"hcall/api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize config
	config.LoadConfig()

	// Initialize database
	database.InitDB()

	// Create Gin router
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	// Setup routes
	routes.SetupRoutes(router)

	// Setup all workers
	workManager := utils.GetWorkerManager()
	go workManager.StartAllWorkers()

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
