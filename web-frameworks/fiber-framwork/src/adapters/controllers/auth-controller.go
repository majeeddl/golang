package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func login(c *fiber.Ctx) error {

	type LoginRequest struct {
		User string `json:"user" form:"user" query:"user" xml:"user" validate:"required"`
		Pass string `json:"pass" form:"pass" query:"pass" xml:"pass" validate:"required"`
	}

	loginInfo := new(LoginRequest)

	if err := c.BodyParser(loginInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request",
		})
	}

	fmt.Println(loginInfo)
	user := loginInfo.User
	pass := loginInfo.Pass

	// fmt.Println(user, pass)

	if user != "admin" || pass != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"user":  user,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func AuthController(app *fiber.App) {
	app.Post("/login", login)
}
