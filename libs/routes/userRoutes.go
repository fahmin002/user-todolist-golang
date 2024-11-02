package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotolist/libs/handlers"
)

func SetupUserRoutes(router fiber.Router) {
	users := router.Group("/users")

	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUserById)
	users.Post("/", handlers.CreateUser)
	users.Put("/:id", handlers.EditUser)
	users.Delete("/:id", handlers.DeleteUser)
}
