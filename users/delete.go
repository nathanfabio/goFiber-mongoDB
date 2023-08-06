package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nathanfabio/goFiber-mongoDB/db"
)

func deleteUser(c *fiber.Ctx) error {
	err := db.DeleteOne("users", c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusNoContent).SendString("")
}