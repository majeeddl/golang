package middlewares

import "github.com/gofiber/fiber/v2"

func OrderMiddleware(app *fiber.App) fiber.Router {

	return app.Use("/orders", func(ctx *fiber.Ctx) error {
		// Do something before

		return ctx.Next()
	})
}
