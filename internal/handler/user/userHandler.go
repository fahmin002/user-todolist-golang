package userHandler

import (
	"fahmin002.github.io/users-api/database"
	model "fahmin002.github.io/users-api/internal/model/user"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	// find all user
	db.Find(&users)

	// if no user is present
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "error", "message": "No user present", "data": users})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Add a uuid to the user
	user.ID = uuid.New()
	// create user
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created User", "data": user})

}
