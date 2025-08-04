package middlewares

import (
	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

func Creator() fiber.Handler {
	return func(c *fiber.Ctx) error {
		permissions := c.Locals("permissions")

		if permissions == nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if permissions == nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		for _, permission := range permissions.([]interface{}) {
			if domain.ToPermissionType(permission) == domain.PermissionCreate {
				return c.Next()
			}
		}

		return c.SendStatus(fiber.StatusUnauthorized)
	}
}
