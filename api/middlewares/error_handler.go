package middlewares

import (
	"net/http"

	"hcall/api/dictionaries"

	"github.com/gin-gonic/gin"
)

// ErrorHandler é um middleware para garantir que todas as respostas de erro também sejam em JSON
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Se houver erros, converta para JSON
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": dictionaries.InternalServerError,
				"status":  false,
			})
		}

		// Manipula rotas não encontradas
		if c.Writer.Status() == http.StatusNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": dictionaries.RouteNotFound,
				"status":  false,
			})
		}
	}
}
