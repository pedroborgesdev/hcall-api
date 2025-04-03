package middlewares

import (
	"regexp"

	"hcall/api/config"
	"hcall/api/logger"
	"hcall/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errors []ValidationError

		// Validate Content-Type
		if c.GetHeader("Content-Type") != "application/json" {
			errors = append(errors, ValidationError{
				Field:   "Content-Type",
				Message: "Content-Type must be application/json",
			})
		}

		// Validate query parameters
		for key, values := range c.Request.URL.Query() {
			if len(values) == 0 {
				errors = append(errors, ValidationError{
					Field:   key,
					Message: "Query parameter cannot be empty",
				})
			}
		}

		// Validate path parameters
		for _, param := range c.Params {
			if param.Value == "" {
				errors = append(errors, ValidationError{
					Field:   param.Key,
					Message: "Path parameter cannot be empty",
				})
			}
		}

		// If there are validation errors, return them
		if len(errors) > 0 {
			utils.SendError(c, utils.CodeValidationError, utils.MsgValidationError, nil)
			return
		}

		c.Next()
	}
}

// ValidatePassword checks if the password meets the security requirements
func ValidatePassword(password string) []string {
	var errors []string

	if len(password) < config.AppConfig.PasswordMinChar {
		errors = append(errors, "Password must be at least 8 characters long")
	}

	if config.AppConfig.PasswordUppercase == "True" {
		upper := regexp.MustCompile(`[A-Z]`)
		if !upper.MatchString(password) {
			errors = append(errors, "Password must contain at least one uppercase letter")
		}
	}

	if config.AppConfig.PasswordLowercase == "True" {
		lower := regexp.MustCompile(`[a-z]`)
		if !lower.MatchString(password) {
			errors = append(errors, "Password must contain at least one lowercase letter")
		}
	}

	if config.AppConfig.PasswordDigits == "True" {
		number := regexp.MustCompile(`[0-9]`)
		if !number.MatchString(password) {
			errors = append(errors, "Password must contain at least one number")
		}
	}

	if config.AppConfig.PasswordSpecial == "True" {
		special := regexp.MustCompile(`[!@#$%^&*]`)
		if !special.MatchString(password) {
			errors = append(errors, "Password must contain at least one special character")
		}
	}

	return errors
}

// ValidateUsername checks if the username meets the requirements
func ValidateUsername(username string) []string {
	var errors []string

	if len(username) < config.AppConfig.UsernameMinChar {
		errors = append(errors, "Username must be at least 6 characters long")
	}

	valid := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !valid.MatchString(username) {
		errors = append(errors, "Username can only contain letters, numbers, and underscores")
	}

	return errors
}

// ValidationMiddleware validates the request body against the provided struct
func ValidationMiddleware(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(v); err != nil {
			logger.Warning("Validation Middleware: Validation failed", map[string]interface{}{
				"error": err.Error(),
			})

			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				utils.SendError(c, utils.CodeValidationError, utils.MsgValidationError, validationErrors)
			} else {
				utils.SendError(c, utils.CodeBadRequest, utils.MsgBadRequest, err)
			}
			c.Abort()
			return
		}

		logger.Debug("Validation Middleware: Validation successful", map[string]interface{}{
			"struct": v,
		})

		c.Next()
	}
}
