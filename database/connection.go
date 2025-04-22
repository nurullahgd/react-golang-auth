package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nurullahgd/react-golang-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	// Upload .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create DSN for database connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Veritabanına bağlan
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanı bağlantısı başarısız:", err)
	}
	db.AutoMigrate(&models.User{})

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Veritabanı bağlantısı başarısız:", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Veritabanı bağlantısını kapatırken hata:", err)
		}
	}()
}
