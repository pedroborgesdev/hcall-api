package utils

import (
	"errors"
	"time"

	"hcall/api/config"
	"hcall/api/logger"
	"hcall/api/models"

	"github.com/golang-jwt/jwt/v5"
)

// Custom claims structure
type JWTClaims struct {
	ID    uint        `json:"id"`
	Email string      `json:"email"`
	Role  models.Role `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token for a user
func GenerateToken(user *models.User) (string, error) {
	logger.Info("Jwt Middleware: Generating token for user", map[string]interface{}{
		"email": user.Email,
		"id":    user.ID,
	})

	// Calculate expiration time
	expirationTime := time.Now().Add(time.Hour * time.Duration(config.AppConfig.JWTExpirationHours))

	claims := &JWTClaims{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))

	if err != nil {
		logger.Error("Jwt middleware: Error generating token", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}

	logger.Info("Jwt middleware: Token generated successfully", map[string]interface{}{
		"email": user.Email,
	})
	return tokenString, nil
}

// ValidateToken validates a JWT token
func ValidateToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil {
		logger.Error("Jwt middleware: Token validation error", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}

	if !token.Valid {
		logger.Error("Jwt middleware: Token is invalid", map[string]interface{}{
			"error": "invalid token",
		})
		return nil, errors.New("invalid token")
	}

	logger.Info("Jwt middleware: Token validated successfully", map[string]interface{}{
		"email": claims.Email,
		"id":    claims.ID,
		"role":  claims.Role,
	})
	return claims, nil
}
