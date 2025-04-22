package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/nurullahgd/react-golang-auth/database"
	"github.com/nurullahgd/react-golang-auth/routes"
)

func main() {

	database.Connect()

	// Initialize a new Fiber app
	app := fiber.New()

	// Setup routes
	routes.Setup(app)
	// Start the server on port 8000

	log.Fatal(app.Listen(":8000"))
}
