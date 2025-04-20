package fibers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// get data and verify
func FormData[T any](c *fiber.Ctx, v *T) (T, error) {
	err := c.BodyParser(&v)
	if err != nil {
		return *v, err
	}
	err = Validate(v)
	if err != nil {
		return *v, err
	}
	return *v, err
}

// get pagination by query
func GetPagination(c *fiber.Ctx) (int, int, int64) {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	size, _ := strconv.Atoi(c.Query("size", "20"))
	return page, size, 0
}

// path variable to int: no error
// tip: this is constraint func, if u cant ensure key can get value you must input defaultValue
func ParamsInt(c *fiber.Ctx, key string, defaultValue ...int) int {
	value, err := c.ParamsInt(key, defaultValue...)
	if err != nil {
		value = -1
	}
	return value
}

// path variable to uint: no error
// tip: this is constraint func, if u cant ensure key can get value you must input defaultValue
func ParamsUInt(c *fiber.Ctx, key string, defaultValue ...int) uint {
	value, err := c.ParamsInt(key, defaultValue...)
	if err != nil {
		value = 0
	}
	return uint(value)
}
