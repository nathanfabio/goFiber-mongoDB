package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathanfabio/goFiber-mongoDB/tags"
	"github.com/nathanfabio/goFiber-mongoDB/tasks"
	"github.com/nathanfabio/goFiber-mongoDB/users"
)

func main() {
	app := fiber.New()

	v1 := app.Group("/v1")
	users.SetRoutes(v1)
	tags.RoutesTags(v1)
	tasks.RoutesTasks(v1)

	app.Listen(":3000")
}