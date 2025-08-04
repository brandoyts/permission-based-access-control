package handlers

import (
	"context"

	"github.com/brandoyts/permission-based-access-control/internal/core/domain"
	"github.com/brandoyts/permission-based-access-control/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	body := new(domain.User)

	err := c.BodyParser(body)
	if err != nil {
		return err
	}

	user, err := uh.service.CreateUser(context.Background(), *body)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (uh *UserHandler) FindUserById(c *fiber.Ctx) error {

	id := c.Params("id")

	user, err := uh.service.FindUserById(context.Background(), id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
