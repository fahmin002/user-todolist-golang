package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gotolist/libs/database"
	"github.com/gotolist/libs/models"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.ConnectDB()
	users := make([]models.User, 0)

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user = new(models.User)
		if err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password_hash,
			&user.Password_salt,
			&user.Created_at,
			&user.Updated_at,
		); err != nil {
			log.Fatal(err)
		}
		users = append([]models.User(users), *user)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data: ": users})
}

func GetUserById(c *fiber.Ctx) error {
	db := database.ConnectDB()
	id := c.Params("id")
	user := new(models.User)
	rows, err := db.Query("SELECT * FROM users where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	if !rows.Next() {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "Error", "message": "user with id :" + id + " not found"})
	} else {
		for rows.Next() {
			if err := rows.Scan(
				&user.Id,
				&user.Username,
				&user.Email,
				&user.Password_hash,
				&user.Password_salt,
				&user.Created_at,
				&user.Updated_at,
			); err != nil {
				log.Fatal(err)
			}
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data: ": user})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.ConnectDB()
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	userId := uuid.New()
	err := db.QueryRow(
		`INSERT INTO users (id, username, email, password_hash, password_salt) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		userId,
		user.Username,
		user.Email,
		user.Password_hash,
		user.Password_salt,
	).Scan(&userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err})
	}
	return c.Status(200).JSON(fiber.Map{"message": "data recorded with id: " + userId.String()})
}

func EditUser(c *fiber.Ctx) error {
	db := database.ConnectDB()
	paramsId := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	err := db.QueryRow(
		`UPDATE users SET username = $1, email = $2, password_hash = $3, password_salt = $4 WHERE id = $5 RETURNING id`,
		user.Username,
		user.Email,
		user.Password_hash,
		user.Password_salt,
		paramsId,
	).Scan(&paramsId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "OK", "message": "user with id: " + paramsId + " updated"})
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.ConnectDB()
	paramsId := c.Params("id")
	res, err := db.Query(
		`DELETE FROM users WHERE id = $1`,
		paramsId,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "OK", "message": "user with id: " + paramsId + " deleted", "Rows Affected": res})
}
