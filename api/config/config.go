package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	UsernameMinChar   int
	PasswordMinChar   int
	PasswordDigits    string
	PasswordSpecial   string
	PasswordUppercase string
	PasswordLowercase string

	Port string

	JWTSecret          string
	JWTExpirationHours int
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	AppConfig = Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "hcall"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		UsernameMinChar:   getEnvInt("USERNAME_MIN_CHAR", 6),
		PasswordMinChar:   getEnvInt("PASSWORD_MIN_CHAR", 8),
		PasswordDigits:    getEnv("PASSWORD_DIGITS", "True"),
		PasswordSpecial:   getEnv("PASSWORD_SPECIAL", "True"),
		PasswordUppercase: getEnv("PASSWORD_UPPERCASE", "True"),
		PasswordLowercase: getEnv("PASSWORD_LOWERCASE", "True"),

		Port: getEnv("PORT", "8080"),

		JWTSecret:          getEnv("JWT_SECRET", "default_jwt_secret_change_this_in_production"),
		JWTExpirationHours: getEnvInt("JWT_EXPIRATION_HOURS", 0),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Invalid value for %s, using default %d: %v", key, defaultValue, err)
		return defaultValue
	}
	return value
}
