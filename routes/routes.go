package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nurullahgd/react-golang-auth/controllers"
)

func Setup(app *fiber.App) {
	// Define your routes here
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// Add more routes as needed

}
