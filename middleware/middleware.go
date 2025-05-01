package middleware

import (
	"web-service/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthentication(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}
