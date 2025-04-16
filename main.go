package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv" // .env dosyasını okuyabilmek için
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Veritabanı bağlantısı için DSN oluştur
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
	// Veritabanı bağlantısını kontrol et
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Veritabanı bağlantısı başarısız:", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Veritabanı bağlantısını kapatırken hata:", err)
		}
	}()

	// Fiber başlat
	app := fiber.New()

	log.Fatal(app.Listen(":8000"))
}
