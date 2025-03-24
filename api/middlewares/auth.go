package middlewares

import (
	"net/http"
	"strings"

	"hcall/api/models"
	"hcall/api/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies the JWT token in the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")

		// Check if Authorization header exists
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"status":  false,
			})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		// Expected format: "Bearer [token]"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"status":  false,
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Validate the token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"status":  false,
			})
			c.Abort()
			return
		}

		// Set the user information in the context for later use
		c.Set("userId", claims.ID)
		c.Set("userEmail", claims.Email)
		c.Set("userRole", claims.Role)

		c.Next()
	}
}

// RoleAuthorization checks if the user has the required role
func RoleAuthorization(allowedRoles ...models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user role from context (set by AuthMiddleware)
		role, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"status":  false,
			})
			c.Abort()
			return
		}

		userRole := role.(models.Role)

		// Check if user role is in allowed roles
		allowed := false
		for _, r := range allowedRoles {
			if userRole == r {
				allowed = true
				break
			}
		}

		// Master role has access to everything
		if userRole == models.MasterRole {
			allowed = true
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "User role not authorized for this endpoint",
				"status":  false,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
