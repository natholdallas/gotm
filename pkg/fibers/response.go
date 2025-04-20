package fibers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natholdallas/templates/fibergorm/pkg/enum"
)

func Status(c *fiber.Ctx, status int) error {
	c.Status(status)
	return nil
}

func StatusAndJSON(c *fiber.Ctx, status int, data any) error {
	c.Status(status)
	return c.JSON(data)
}

func Err(value any, statis ...int) *fiber.Error {
	msg := enum.Unknown
	code := fiber.StatusBadRequest
	if str, ok := value.(string); ok {
		msg = str
	} else if err, ok := value.(error); ok {
		msg = err.Error()
	}
	return &fiber.Error{Code: code, Message: msg}
}
