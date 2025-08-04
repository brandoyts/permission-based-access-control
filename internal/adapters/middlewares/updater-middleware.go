package middlewares

import (
	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

func toPermissionType(permission string) domain.PermissionType {
	return domain.PermissionType(permission)
}

func Updater() fiber.Handler {
	return func(c *fiber.Ctx) error {
		permission := c.Get("permission")

		permissionType := toPermissionType(permission)

		if permissionType != domain.PermissionUpdate {
			return c.Status(fiber.StatusUnauthorized).SendString("updated only")
		}

		return c.Next()
	}
}
