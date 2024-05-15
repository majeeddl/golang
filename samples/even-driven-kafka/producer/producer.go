package main

import "github.com/gofiber/fiber/v2"

type Comment struct {
	Comment string `from:"text" json:"text"`
}

func main() {

	app := fiber.New()
	api := app.Group("/api/v1")

	api.Post("/comment", createComment)

	app.Listen(":3000")
}

func createComment(c *fiber.Ctx) error {

	cmt := new(Comment)

}
