package middlewares

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func NewAuthMiddleware(secret string) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		// SigningMethod: "HS256",
		SuccessHandler: func(c *fiber.Ctx) error {
			// user := c.Locals("user").(*jwt.Token)
			// claims := user.Claims.(jwt.MapClaims)
			// c.Locals("user", claims)
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}

func ProtectedMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	fmt.Println(user)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["user"].(string)
	// return c.SendString("Welcome " + username)
	c.Locals("user", username)
	return c.Next()
}
