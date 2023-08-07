package tags

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathanfabio/goFiber-mongoDB/db"
)

func getAll(c *fiber.Ctx) error {

	var documents []Tags

	err := db.Find("tags", nil, &documents)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(documents)
}

func getOne(c *fiber.Ctx) error {
	var document Tags

	err := db.FindOne("tags", c.Params("id"), &document)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(document)
}