package tasks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathanfabio/goFiber-mongoDB/db"
)

func getAll(c *fiber.Ctx) error {

	var documents []Tasks

	err := db.Find("tasks", nil, &documents)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(documents)
}

func getOne(c *fiber.Ctx) error {
	var document Tasks

	err := db.FindOne("tasks", c.Params("id"), &document)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(document)
}