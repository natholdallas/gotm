package handler

import (
	"github.com/gofiber/fiber/v2"
)

func TestHandler(c *fiber.Ctx) error {
	return c.JSON(nil)
}
