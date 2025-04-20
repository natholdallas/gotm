package mid

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natholdallas/templates/fibergorm/pkg/enum"
	"github.com/natholdallas/templates/fibergorm/pkg/fibers"
)

func IsAdmin(c *fiber.Ctx) error {
	claims, err := FindClaims(c)
	if err != nil {
		return err
	}
	if !claims.IsAdmin {
		return fibers.Err(enum.IsAdmin, fiber.StatusForbidden)
	}
	return c.Next()
}

func NoGoogle(c *fiber.Ctx) error {
	claims, err := FindClaims(c)
	if err != nil {
		return err
	}
	if claims.IsGoogleUser {
		return fibers.Err(enum.NoGoogle, fiber.StatusForbidden)
	}
	return c.Next()
}
