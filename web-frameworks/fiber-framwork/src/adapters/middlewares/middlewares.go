package middlewares

import "github.com/gofiber/fiber/v2"

func tenantMiddleware(app *fiber.App) fiber.Router {

	return app.Use("/*", func(ctx *fiber.Ctx) error {
		// Do something before

		tenantId := ctx.Get("x-tenantId")

		if tenantId != "" {
			// return ctx.Status(400).JSON("Tenant Id is required")
		}

		ctx.Locals("tenantId", tenantId)

		return ctx.Next()
	})
}

func SetRoutesMiddlewares(app *fiber.App) error {

	tenantMiddleware(app)
	return nil

}
