package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotolist/libs/router"
)

func main() {
	app := fiber.New()

	router.SetupRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hello World!")
		return err
	})

	app.Listen(":3000")
}
