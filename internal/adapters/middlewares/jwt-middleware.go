package middlewares

import (
	"fmt"
	"strings"

	jwtProvider "github.com/brandoyts/permission-based-access-control/internal/adapters/jwt-provider"
	"github.com/gofiber/fiber/v2"
)

const authorizationHeader = "Authorization"
const bearerTokenLength = 2
const bearerKeyword = "bearer"
const permissions = "permissions"

func JwtAuth(jwtProvider *jwtProvider.JwtProvider) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get(authorizationHeader)

		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != bearerTokenLength {
			fmt.Println("no")
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if bearerToken[0] != bearerKeyword {
			fmt.Println("arrrrr")
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		jwtToken := bearerToken[1]

		tokenClaims, err := jwtProvider.DecodeToken(jwtToken)
		if err != nil {
			fmt.Println(err, jwtToken)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Locals(permissions, tokenClaims[permissions])

		return c.Next()
	}
}
