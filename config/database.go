package config

import (
	"BE_Go/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	loadEnv()

	// Check STATUS to determine environment
	status := os.Getenv("STATUS")
	if status == "prod" {
		// Use PostgreSQL for production
		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			panic("DB_URL must be set when STATUS=prod")
		}
		DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	} else {
		// Use SQLite for development (default)
		DB, err = gorm.Open(sqlite.Open("houses.db"), &gorm.Config{})
	}

	if err != nil {
		panic("failed to connect to database")
	}

	// AutoMigrate the schema
	DB.AutoMigrate(
		&models.House{},
		&models.Admin{},
		&models.Captain{},
		&models.Member{},
		&models.Announcement{},
		&models.PointTransaction{},
	)
}
