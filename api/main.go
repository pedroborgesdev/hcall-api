package main

import (
	"fmt"
	"log"
	"net/http"

	"hcall/api/config"
	"hcall/api/database"
	"hcall/api/dictionaries"
	"hcall/api/middlewares"
	"hcall/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Carrega as variáveis de ambiente
	config.LoadConfig()

	// Conecta ao banco de dados
	database.ConnectDatabase()
	database.MigrateDatabase()

	// Inicializa o router Gin
	router := gin.Default()

	// Adiciona o middleware para garantir respostas em JSON
	router.Use(middlewares.JSONResponseMiddleware())

	// Adiciona o middleware de tratamento de erros
	router.Use(middlewares.ErrorHandler())

	// Configura o CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))

	// Configura as rotas
	routes.SetupRoutes(router)

	// Handler para rotas não encontradas
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": dictionaries.RouteNotFound,
			"status":  false,
		})
	})

	// Inicia o servidor
	port := config.AppConfig.Port
	log.Printf("Server is starting on http://localhost:%s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
