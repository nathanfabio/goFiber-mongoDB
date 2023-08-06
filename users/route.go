package users

import "github.com/gofiber/fiber/v2"

//SetRoutes sets up the routes for the API.
func SetRoutes(r fiber.Router) {
	users := r.Group("/users")

	users.Post("/", addUser)
	
	users.Get("/", getAll)
	users.Get("/:id", getOne)

	users.Put("/:id", updateUser)

	users.Delete("/:id", deleteUser)
}