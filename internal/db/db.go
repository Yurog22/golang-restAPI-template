package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"template/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSL_MODE"),
		os.Getenv("SSL_ROOT_CERTIFICATE"),
	)
	database, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database")
	}

	database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := database.AutoMigrate(&models.Example{}); err != nil {
		panic("Failed to automigrate")
	}

	DB = database
}
