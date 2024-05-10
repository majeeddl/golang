package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/before-middleware", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, Fiber! Before Middleware")
	})

	app.Use(func(c *fiber.Ctx) error {

		fmt.Println("First middleware")

		return c.Next()
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, Fiber!")
	})

	app.Listen(":3000")
}
