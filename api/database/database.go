package database

import (
	"fmt"

	"hcall/api/config"
	"hcall/api/logger"
	"hcall/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and runs migrations
func InitDB() {
	// Get database configuration from environment variables
	host := config.AppConfig.DBHost
	port := config.AppConfig.DBPort
	user := config.AppConfig.DBUser
	password := config.AppConfig.DBPassword
	dbname := config.AppConfig.DBName
	sslmode := config.AppConfig.DBSSLMode

	// Create connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Database: Failed to connect to database:", map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Get the underlying SQL DB
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("Database: Failed to get database instance:", map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Run migrations
	err = runMigrations(db)
	if err != nil {
		logger.Fatal("Database: Failed to migrate database:", map[string]interface{}{
			"error": err.Error(),
		})
	}

	DB = db
	logger.Info("Database: Connected to database successfully", nil)
}

// runMigrations runs all database migrations
func runMigrations(db *gorm.DB) error {
	// Drop the base64 column if it exists (this is safe because we're starting fresh)
	// _ = db.Exec(`ALTER TABLE "images" DROP COLUMN IF EXISTS "base64"`).Error

	// Run GORM migrations to create tables and add the base64 column properly
	err := db.AutoMigrate(
		&models.User{},
		&models.Ticket{},
		&models.Counters{},
		&models.Image{},
		&models.TicketHistory{},
	)
	if err != nil {
		return err
	}

	return nil
}
