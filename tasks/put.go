package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nathanfabio/goFiber-mongoDB/db"
	"github.com/nathanfabio/goFiber-mongoDB/tags"
)

func updateTasks(c *fiber.Ctx) error {
	body := new(Tasks)

	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	var previous Tasks
	err := db.FindOne("tasks", c.Params("id"), &previous)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	var result Tasks
	err = db.UpdateOne("tasks", c.Params("id"), body, &result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = updateTagsTask(c.Params("id"), previous.Tags, result.Tags)
	if err != nil {
		err = tags.RemoveTask(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		err = db.DeleteOne("tasks", c.Params("id"))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		_, err = db.Insert("tasks", &result)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		err = tags.AddTask(result.ID.Hex(), result.Tags)
		if err != nil {
			db.DeleteOne("tasks", c.Params("id"))
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
	}

	return c.JSON(result)
}

func updateTagsTask(id string, oldt, newt []string) error {
	moldl := make(map[string]int, len(oldt))

	for k, v := range oldt {
		moldl[v] = k
	}
	var diff []string
	for _, v := range newt {
		if _, key := moldl[v]; !key {
			diff = append(diff, v)
		} else {
			delete(moldl, v)
		}
	}

	if len(diff) > 0 {
		err := tags.AddTask(id, diff)
		if err != nil {
			return err
		}
	}

	if len(moldl) > 0 {
		dt := make([]string, 0, len(moldl))
		for k := range moldl {
		dt= append(dt, k)	
		}

		err := tags.RemoveTask(id, dt...)
		if err != nil {
			return err
		}
	}

	return nil
}