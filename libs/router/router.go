package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gotolist/libs/routes"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())

	if api != nil {
		routes.SetupUserRoutes(api)
	}
}
