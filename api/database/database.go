package database

import (
	"fmt"
	"log"

	"hcall/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and runs migrations
func InitDB() {
	// PostgreSQL connection configuration
	const (
		host     = "localhost"
		port     = "5432"
		user     = "postgres"
		password = "s1ea021274"
		dbname   = "hcall"
		sslmode  = "disable"
	)

	// Create connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Get the underlying SQL DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Run migrations
	err = runMigrations(db)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	log.Println("Connected to database successfully")
}

// runMigrations runs all database migrations
func runMigrations(db *gorm.DB) error {
	// Drop the base64 column if it exists (this is safe because we're starting fresh)
	// _ = db.Exec(`ALTER TABLE "images" DROP COLUMN IF EXISTS "base64"`).Error

	// Run GORM migrations to create tables and add the base64 column properly
	err := db.AutoMigrate(
		&models.User{},
		&models.Ticket{},
		&models.Image{},
		&models.TicketHistory{},
	)
	if err != nil {
		return err
	}

	return nil
}
