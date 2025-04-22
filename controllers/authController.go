package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nurullahgd/react-golang-auth/helpers"
	"github.com/nurullahgd/react-golang-auth/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c fiber.Ctx) error {
	var data map[string]string

	if err := c.Bind().Body(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 15)

	user := models.User{
		Id:       helpers.GenerateUUID(8),
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	return c.JSON(user)
}
