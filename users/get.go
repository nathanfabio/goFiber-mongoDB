package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathanfabio/goFiber-mongoDB/db"
)

func getAll(c *fiber.Ctx) error {
	var documents []User

	err := db.Find("users", nil, &documents)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(documents)
}

func getOne(c *fiber.Ctx) error {
	var document User

	err := db.FindOne("users", c.Params("id"), &document)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(document)
}