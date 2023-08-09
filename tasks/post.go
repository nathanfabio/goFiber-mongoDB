package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nathanfabio/goFiber-mongoDB/db"
	"github.com/nathanfabio/goFiber-mongoDB/tags"
)

func addTask(c *fiber.Ctx) error {
	body := new(Tasks)

	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	id, err := db.Insert("tasks", body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	body.ID = id

	err = tags.AddTask(body.ID.Hex(), body.Tags)
	if err != nil {
		db.DeleteOne("tasks", body.ID.Hex())
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(body)
}