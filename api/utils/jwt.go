package utils

import (
	"errors"
	"log"
	"time"

	"hcall/api/config"
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
	log.Printf("Generating token for user %s (ID: %d)", user.Email, user.ID)

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
		log.Printf("Error generating token: %v", err)
		return "", err
	}

	log.Printf("Token generated successfully for user %s", user.Email)
	return tokenString, nil
}

// ValidateToken validates a JWT token
func ValidateToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil {
		log.Printf("Token validation error: %v", err)
		return nil, err
	}

	if !token.Valid {
		log.Printf("Token is invalid")
		return nil, errors.New("invalid token")
	}

	log.Printf("Token validated successfully for user %s (ID: %d) with role %s", claims.Email, claims.ID, claims.Role)
	return claims, nil
}
