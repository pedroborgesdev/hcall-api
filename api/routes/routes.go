package routes

import (
	"hcall/api/controllers"
	"hcall/api/middlewares"
	"hcall/api/models"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all routes for the API
func SetupRoutes(router *gin.Engine) {
	// Create controllers
	authController := controllers.NewAuthController()
	userController := controllers.NewUserController()
	ticketController := controllers.NewTicketController()

	// Apply CORS middleware to all routes
	router.Use(middlewares.CORSMiddleware())

	// Grupo de rotas com prefixo '/api'
	api := router.Group("/api")
	{
		// Rotas de autenticação - públicas
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/enter", authController.Login)
		}

		// Rotas para Master - públicas
		master := api.Group("/master")
		{
			master.POST("/create", authController.CreateMaster)
			master.POST("/delete", authController.DeleteMaster)
		}

		// Rotas protegidas - requerem autenticação
		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware())

		// Rotas de usuário - apenas admin e master
		user := protected.Group("/user")
		user.Use(middlewares.RoleAuthorization(models.AdminRole, models.MasterRole))
		{
			user.GET("/fetch", userController.GetUsers)
			user.POST("/create", userController.CreateUser)
			user.POST("/delete", userController.DeleteUser)
		}

		// Rotas de tickets
		ticket := protected.Group("/ticket")
		{
			// Rotas acessíveis a usuários, admins e masters
			ticket.POST("/create", ticketController.CreateTicket)
			ticket.POST("/remove", ticketController.DeleteTicket)
			ticket.GET("/count", ticketController.CountTicket)

			// Rotas acessíveis apenas a admins e masters
			authTicket := ticket.Group("/")
			authTicket.Use(middlewares.RoleAuthorization(models.AdminRole, models.MasterRole))
			{
				authTicket.GET("/fetch", ticketController.GetTickets)
				authTicket.GET("/info", ticketController.GetTicketDetails)
				authTicket.POST("/edit", ticketController.UpdateTicketStatus)
				authTicket.POST("/update", ticketController.UpdateTicketHistory)
			}
		}
	}

	// Start the ticket worker

}
