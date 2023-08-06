package tags

import "github.com/gofiber/fiber/v2"

func RoutesTags(r fiber.Router) {
	tags := r.Group("/tags")

	tags.Post("/", addTag)

	tags.Get("/", getAll)
	tags.Get("/:id", getOne)

	tags.Put("/:id", updateTags)

	tags.Delete("/:id", deleteTags)
}