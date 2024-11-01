package router

import (
	userRoutes "fahmin002.github.io/users-api/internal/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	if api != nil {
		userRoutes.SetupUserRoutes(api)
	}
}
