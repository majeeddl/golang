package controllers

import (
	"fiberframework/src/domain/dto"
	"fiberframework/src/domain/structs"
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
func getOrders(config structs.ControllerConfig) fiber.Router {
	return config.App.Get("/orders", func(ctx *fiber.Ctx) error {

		result, err := config.DataService.OrderRepository().FindAll()

		result2, err := config.DataService.OrderRepository().Test()

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
func getOrderById(config structs.ControllerConfig) fiber.Router {
	return config.App.Get("/orders/:id", func(ctx *fiber.Ctx) error {

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
func createOrder(config structs.ControllerConfig) fiber.Router {
	return config.App.Post("/orders", func(ctx *fiber.Ctx) error {

		orderUseCase := usecases.OrderUseCase{}

		orderUseCase.CreateOrder()

		var request dto.CreateOrderDTO

		err := ctx.BodyParser(&request)

		if err != nil {
			return ctx.Status(400).JSON("Error parsing request")
		}

		if errs := config.Validator.Validate(config.Validator.Validator, request); errs != nil {
			// errorMessages := make([]string, len(errs))

			// for i, err := range errs {
			// 	errorMessages[i] = fmt.Sprintf("Field: %s, Tag: %s, Value: %v", err.Field, err.Tag, err.Value)
			// }

			return ctx.Status(400).JSON(errs)
		}

		return ctx.SendString("Create Order")
	})
}

func OrdersController(config structs.ControllerConfig) {
	getOrders(config)
	getOrderById(config)
	createOrder(config)
}
