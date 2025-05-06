package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/nurullahgd/react-golang-auth/database"
	"github.com/nurullahgd/react-golang-auth/routes"
)

func main() {

	database.Connect()

	// Initialize a new Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // React uygulamanızın çalıştığı URL
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	// Setup routes
	routes.Setup(app)
	// Start the server on port 8000

	log.Fatal(app.Listen(":8000"))
}
