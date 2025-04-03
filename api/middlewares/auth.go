package middlewares

import (
	"fmt"
	"strings"

	"hcall/api/logger"
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
			logger.Warning("Auth Middleware: Missing Authorization header", map[string]interface{}{
				"ip": c.ClientIP(),
			})
			utils.SendError(c, utils.CodeUnauthorized, utils.MsgUnauthorized, nil)
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		// Expected format: "Bearer [token]"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logger.Warning("Auth Middleware: Invalid Authorization header format", map[string]interface{}{
				"ip": c.ClientIP(),
			})
			utils.SendError(c, utils.CodeUnauthorized, utils.MsgUnauthorized, nil)
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Validate the token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			logger.Error("Auth Middleware: Token validation failed", map[string]interface{}{
				"ip":    c.ClientIP(),
				"error": err.Error(),
			})
			utils.SendError(c, utils.CodeUnauthorized, utils.MsgUnauthorized, err)
			c.Abort()
			return
		}

		logger.Info("Auth Middleware: Token validated successfully", map[string]interface{}{
			"user_id": claims.ID,
			"email":   claims.Email,
			"role":    claims.Role,
		})

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
			logger.Warning("Auth Middleware: User role not found in context", map[string]interface{}{
				"ip": c.ClientIP(),
			})
			utils.SendError(c, utils.CodeUnauthorized, utils.MsgUnauthorized, nil)
			c.Abort()
			return
		}

		userRole, ok := role.(models.Role)
		if !ok {
			logger.Error("Auth Middleware: Failed to convert role to models.Role", map[string]interface{}{
				"ip":   c.ClientIP(),
				"role": role,
				"type": fmt.Sprintf("%T", role),
			})
			utils.SendError(c, utils.CodeUnauthorized, utils.MsgUnauthorized, nil)
			c.Abort()
			return
		}

		logger.Info("Auth Middleware: Checking role authorization", map[string]interface{}{
			"user_role":     userRole,
			"allowed_roles": allowedRoles,
		})

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
			logger.Warning("Auth Middleware: Role not authorized", map[string]interface{}{
				"user_role":     userRole,
				"allowed_roles": allowedRoles,
			})
			utils.SendError(c, utils.CodeForbidden, utils.MsgForbidden, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
