package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	DBMaxIdleConns int
	DBMaxOpenConns int
	DBConnTimeout  int

	UsernameMinChar   int
	PasswordMinChar   int
	PasswordDigits    string
	PasswordSpecial   string
	PasswordUppercase string
	PasswordLowercase string

	WorkerTicketLooptime    int
	WorkerTicketRemoveAfter int
	WorkerTicketStatus      string

	Port string

	JWTSecret          string
	JWTExpirationHours int

	// Rate Limiting
	RateLimitRequests int
	RateLimitWindow   int

	Debug   bool
	GINMode bool
}

var AppConfig Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Environment file not found, using environment variables")
	}

	AppConfig = Config{
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "hcall"),
		DBSSLMode:      getEnv("DB_SSLMODE", "disable"),
		DBMaxIdleConns: getEnvInt("DB_MAX_IDLE_CONNS", 10),
		DBMaxOpenConns: getEnvInt("DB_MAX_OPEN_CONNS", 100),
		DBConnTimeout:  getEnvInt("DB_CONN_TIMEOUT", 5),

		UsernameMinChar:   getEnvInt("USERNAME_MIN_CHAR", 6),
		PasswordMinChar:   getEnvInt("PASSWORD_MIN_CHAR", 8),
		PasswordDigits:    getEnv("PASSWORD_DIGITS", "True"),
		PasswordSpecial:   getEnv("PASSWORD_SPECIAL", "True"),
		PasswordUppercase: getEnv("PASSWORD_UPPERCASE", "True"),
		PasswordLowercase: getEnv("PASSWORD_LOWERCASE", "True"),

		WorkerTicketLooptime:    getEnvInt("WORKER_TICKET_LOOPTIME", 30),
		WorkerTicketRemoveAfter: getEnvInt("WORKER_TICKET_REMOVE_AFTER", 10),
		WorkerTicketStatus:      getEnv("WORKER_TICKET_REMOVE_STATUS", "conclued"),

		Port: getEnv("PORT", "8080"),

		JWTSecret:          getEnv("JWT_SECRET", "default_jwt_secret_change_this_in_production"),
		JWTExpirationHours: getEnvInt("JWT_EXPIRATION_HOURS", 24),

		RateLimitRequests: getEnvInt("RATE_LIMIT_REQUESTS", 100),
		RateLimitWindow:   getEnvInt("RATE_LIMIT_WINDOW", 60),

		Debug:   getEnvBool("DEBUG", true),
		GINMode: getEnvBool("GIN_MODE", true),
	}

	return AppConfig.Validate()
}

func (c *Config) Validate() error {
	if c.DBHost == "" {
		return errors.New("DB_HOST is required")
	}
	if c.DBPort == "" {
		return errors.New("DB_PORT is required")
	}
	if c.DBUser == "" {
		return errors.New("DB_USER is required")
	}
	if c.DBPassword == "" {
		return errors.New("DB_PASSWORD is required")
	}
	if c.DBName == "" {
		return errors.New("DB_NAME is required")
	}
	if c.JWTSecret == "" {
		return errors.New("JWT_SECRET is required")
	}
	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}
