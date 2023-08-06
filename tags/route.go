package tags

import "github.com/gofiber/fiber/v2"

func RoutesTags(r fiber.Router) {
	tags := r.Group("/tags")

	tags.Post("/", addTag)
}