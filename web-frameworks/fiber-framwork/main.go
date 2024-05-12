package main

import (
	"fiberframework/src/adapters/controllers"
	"fiberframework/src/adapters/middlewares"
	"fiberframework/src/adapters/sockets"

	"fiberframework/src/adapters/validation"
	"fiberframework/src/configuration"
	"fmt"
	"log"

	_ "fiberframework/docs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/recover"
	// "github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

var validate = validator.New()

// @title			Fiber Framework API
// @version		1.0
// @description	This is a sample swagger for Fiber
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// app.Use(cache.New(cache.Config{
	// 	Next: func(c *fiber.Ctx) bool {
	// 		return strings.Contains(c.Route().Path, "/ws")
	// 	},
	// }))
	// app.Get("/swagger/*", swagger.HandlerDefault)

	customValidator := &validation.Validation{
		Validator: validate,
	}

	// err = validation.SetRoutesValidation(customValidator)

	// if err != nil {
	// 	log.Fatal("Error setting validation rules")
	// }

	// app.Use(recover.New())
	// app.Use(cors.New())

	middlewares.SetRoutesMiddlewares(app)

	controllers.OrdersController(app, customValidator)

	sockets.InitializeSockets(app)

	app.Listen(fmt.Sprintf(":%s", configuration.GetPort()))
}
