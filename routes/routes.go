package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nurullahgd/react-golang-auth/controllers"
)

func Setup(app *fiber.App) {
	// Define your routes here
	app.Post("/api/register", func(c fiber.Ctx) error {
		return controllers.Register(c)
	})

	// Add more routes as needed

}
