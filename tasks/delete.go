package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nathanfabio/goFiber-mongoDB/db"
)

func deleteTasks(c *fiber.Ctx) error {
	err := db.DeleteOne("tasks", c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusNoContent).SendString("")
}