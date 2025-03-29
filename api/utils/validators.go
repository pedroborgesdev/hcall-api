package utils

import (
	"fmt"
	"hcall/api/config"
	"strings"
)

func ValidateCredentials(email, password, username string) error {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return fmt.Errorf("Invalid email format")
	}

	if len(username) < config.AppConfig.UsernameMinChar {
		return fmt.Errorf("Username must be at least %d characters long", config.AppConfig.UsernameMinChar)
	}

	if len(password) < config.AppConfig.PasswordMinChar {
		return fmt.Errorf("Password must be at least %d characters long", config.AppConfig.PasswordMinChar)
	}

	if config.AppConfig.PasswordUppercase == "True" {
		if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			return fmt.Errorf("Password must have at least one uppercase character")
		}
	}

	if config.AppConfig.PasswordLowercase == "True" {
		if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
			return fmt.Errorf("Password must have at least one lowercase character")
		}
	}

	if config.AppConfig.PasswordSpecial == "True" {
		if !strings.ContainsAny(password, "!@#$%^&*()-_=+[]{}|;:',.<>/?") {
			return fmt.Errorf("Password must have at least one special character")
		}
	}

	return nil
}
