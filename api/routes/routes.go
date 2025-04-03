package routes

import (
	"hcall/api/controllers"
	"hcall/api/middlewares"
	"hcall/api/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// CORS aplicado GLOBALMENTE (inclui /health e todas as rotas /api)

	// Controllers
	authController := controllers.NewAuthController()
	userController := controllers.NewUserController()
	ticketController := controllers.NewTicketController()

	// Rota de health check (pública)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Grupo /api
	api := router.Group("/api")
	{
		// Rotas PÚBLICAS (sem autenticação)
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/enter", authController.Login) // Rota de login
		}

		// Rotas MASTER (protegidas por outro mecanismo, não por JWT)
		master := api.Group("/master")
		{
			master.POST("/create", authController.CreateMaster)
			master.POST("/delete", authController.DeleteMaster)
		}

		// Rotas PROTEGIDAS (exigem JWT)
		protected := api.Group("/")
		protected.Use(middlewares.AuthMiddleware()) // Middleware de JWT
		{
			// Rotas de usuário (apenas admin e master)
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
				ticket.POST("/create", ticketController.CreateTicket)
				ticket.POST("/remove", ticketController.DeleteTicket)
				ticket.GET("/count", ticketController.CountTicket)

				// Subgrupo com permissão adicional (admin/master)
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
	}
}
