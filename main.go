package main

import (
	"encoding/json"

	"fahmin002.github.io/users-api/database"
	"fahmin002.github.io/users-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start Fiber App
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Send a string for GET Calls to "/" endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	// connect to database
	database.ConnectDB()

	// setup router
	router.SetupRoutes(app)

	app.Listen("localhost:3000")
}
