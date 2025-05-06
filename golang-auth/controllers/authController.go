package controllers

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nurullahgd/react-golang-auth/database"
	"github.com/nurullahgd/react-golang-auth/helpers"
	"github.com/nurullahgd/react-golang-auth/models"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "Secret"

func Register(c fiber.Ctx) error {
	var data map[string]string

	if err := c.Bind().Body(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 15)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not hash password",
		})
	}

	user := models.User{
		Id:       helpers.GenerateUUID(8),
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}

	return c.JSON(user)
}

func Login(c fiber.Ctx) error {
	var data map[string]string

	if err := c.Bind().Body(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}
	var user models.User

	database.DB.Where("email=?", data["email"]).First(&user)

	if user.Id == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	claims := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
		"exp":   jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return err
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Login successful",
	})

}
func User(c fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	claims := token.Claims.(*jwt.MapClaims)

	var user models.User

	database.DB.Where("email=?", (*claims)["email"]).First(&user)

	return c.JSON(claims)

}

func Logout(c fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
