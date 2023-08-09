package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nathanfabio/goFiber-mongoDB/db"
	"github.com/nathanfabio/goFiber-mongoDB/tags"
)

func deleteTasks(c *fiber.Ctx) error {
	err := tags.RemoveTask(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = db.DeleteOne("tasks", c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusNoContent).SendString("")
}