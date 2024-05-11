package controllers

import (
	"fiberframework/src/adapters/validation"
	"fiberframework/src/domain/dto"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Get order by ID
// @Description	get string by ID
// @Tags			orders
// @Accept			json
// @Produce		json
// @Param			id			path		int		true	"Order ID"
// @Param			x-tenantId	header		string	true	"Tenant ID"	default("favana")
// @Success		200			{object}	string
// @Router			/orders/{id} [get]
func getOrderById(app *fiber.App) fiber.Router {
	return app.Get("/orders/:id", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Get Order By Id :" + ctx.Params("id") + ", tenantId: " + ctx.Locals("tenantId").(string))
	})
}

// @Summary		Create an order
// @Description	Create an order
// @Tags			orders
// @Accept			json
// @Produce		json
// @Param			dto.createOrderDTO	body		dto.CreateOrderDTO	true	"Order object that needs to be created"
// @Param			x-tenantId			header		string				true	"Tenant ID"	default("favana")
// @Success		200					{object}	string
// @Router			/orders [post]
func createOrder(app *fiber.App, validator *validation.Validation) fiber.Router {
	return app.Post("/orders", func(ctx *fiber.Ctx) error {

		var request dto.CreateOrderDTO

		err := ctx.BodyParser(&request)

		if err != nil {
			return ctx.Status(400).JSON("Error parsing request")
		}

		if errs := validator.Validate(validator.Validator, request); errs != nil {
			// errorMessages := make([]string, len(errs))

			// for i, err := range errs {
			// 	errorMessages[i] = fmt.Sprintf("Field: %s, Tag: %s, Value: %v", err.Field, err.Tag, err.Value)
			// }

			return ctx.Status(400).JSON(errs)
		}

		return ctx.SendString("Create Order")
	})
}

func OrdersController(app *fiber.App, validator *validation.Validation) {
	getOrderById(app)
	createOrder(app, validator)
}
