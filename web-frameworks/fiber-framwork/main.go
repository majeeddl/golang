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
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/golang-jwt/jwt/v5"

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
	app.Use(cors.New())
	app.Use(csrf.New())

	controllers.AuthController(app)

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

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	// Restricted Routes
	app.Get("/restricted", restricted)

	sockets.InitializeSockets(app)

	app.Listen(fmt.Sprintf(":%s", configuration.GetPort()))
}

func restricted(c *fiber.Ctx) error {
	fmt.Println("Restricted")
	fmt.Println(c.Locals("user"))
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	username := claims["user"].(string)
	// name := "admin"
	return c.SendString("Welcome " + username)
}
