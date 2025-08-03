package main

import (
	"github.com/brandoyts/permission-based-access-control/internal/adapters/handlers"
	"github.com/brandoyts/permission-based-access-control/internal/adapters/persistence/memory"
	"github.com/brandoyts/permission-based-access-control/internal/core/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())

	userRepository := memory.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("success")
	})

	userRouter := app.Group("users")
	userRouter.Post("/create", userHandler.CreateUser)
	userRouter.Get("/:id", userHandler.FindUserById)

	app.Listen(":9094")
}
