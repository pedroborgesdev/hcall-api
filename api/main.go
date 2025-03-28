package main

import (
	"log"

	"hcall/api/config"
	"hcall/api/database"
	"hcall/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize config
	config.LoadConfig()

	// Initialize database
	database.InitDB()

	// Create Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
