package tasks

import "github.com/gofiber/fiber/v2"

func RoutesTasks(r fiber.Router) {
	tasks := r.Group("/tasks")

	tasks.Post("/", addTask)

	tasks.Get("/", getAll)
	tasks.Get("/:id", getOne)

	tasks.Put("/:id", updateTags)
}