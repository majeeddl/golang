package controllers

import (
	"fiberframework/src/adapters/validation"
	"fiberframework/src/domain/dto"
	"fiberframework/src/domain/interfaces"
	usecases "fiberframework/src/use-cases/orders"
	"fmt"

	// "github.com/gofiber/contrib/socketio"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Get orders
// @Description	get orders
// @Tags			orders
// @Accept			json
// @Produce		json
// @Param			x-tenantId	header		string	true	"Tenant ID"	default("favana")
// @Success		200			{object}	string
// @Router			/orders [get]
func getOrders(app *fiber.App, dataservice interfaces.IDataService) fiber.Router {
	return app.Get("/orders", func(ctx *fiber.Ctx) error {

		result, err := dataservice.OrderRepository().FindAll()

		result2, err := dataservice.OrderRepository().Test()

		fmt.Println(result2)

		if err != nil {
			return ctx.Status(500).JSON(err)
		}

		return ctx.JSON(result)
		// socketio.Broadcast([]byte("Hello World"))
		// return ctx.SendString("Get Order By Id :" + ctx.Params("id") + ", tenantId: " + ctx.Locals("tenantId").(string))
	})
}

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

		socketio.Broadcast([]byte("Hello World"))
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

		orderUseCase := usecases.OrderUseCase{}

		orderUseCase.CreateOrder()

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

func OrdersController(app *fiber.App, dataservice interfaces.IDataService, validator *validation.Validation) {
	getOrders(app, dataservice)
	getOrderById(app)
	createOrder(app, validator)
}
