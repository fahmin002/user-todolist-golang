package userRoutes

import (
	userHandler "fahmin002.github.io/users-api/internal/handler/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.GetUsers)

}
