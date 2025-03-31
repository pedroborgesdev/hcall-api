package utils

import (
	"fmt"
	"hcall/api/config"
	"hcall/api/dictionaries"
	"strings"
)

func ValidateCredentials(email, password, username string) error {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return fmt.Errorf(dictionaries.InvalidEmailFormat)
	}

	if len(username) < config.AppConfig.UsernameMinChar {
		return fmt.Errorf(dictionaries.UsernameTooShort, config.AppConfig.UsernameMinChar)
	}

	if len(password) < config.AppConfig.PasswordMinChar {
		return fmt.Errorf(dictionaries.PasswordTooShort, config.AppConfig.PasswordMinChar)
	}

	if config.AppConfig.PasswordUppercase == "True" {
		if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			return fmt.Errorf(dictionaries.PasswordUppercase)
		}
	}

	if config.AppConfig.PasswordLowercase == "True" {
		if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
			return fmt.Errorf(dictionaries.PasswordLowercase)
		}
	}

	if config.AppConfig.PasswordSpecial == "True" {
		if !strings.ContainsAny(password, "!@#$%^&*()-_=+[]{}|;:',.<>/?") {
			return fmt.Errorf(dictionaries.PasswordSpecial)
		}
	}

	return nil
}
