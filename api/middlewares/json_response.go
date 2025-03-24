package middlewares

import (
	"github.com/gin-gonic/gin"
)

// JSONResponseMiddleware ensures all responses are in JSON format
func JSONResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}
