package handlers

import (
	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
	"github.com/brandoyts/permission-based-access-control/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service ports.AuthService
}

func NewAuthHandler(service ports.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	body := new(domain.User)

	c.BodyParser(body)

	token, err := ah.service.Login(body.Email, body.Password)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.JSON(token)
}
