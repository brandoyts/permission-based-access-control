package main

import (
	"github.com/brandoyts/permission-based-access-control/internal/adapters/handlers"
	jwtProvider "github.com/brandoyts/permission-based-access-control/internal/adapters/jwt-provider"
	"github.com/brandoyts/permission-based-access-control/internal/adapters/middlewares"
	"github.com/brandoyts/permission-based-access-control/internal/adapters/persistence/memory"
	"github.com/brandoyts/permission-based-access-control/internal/core/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())

	jwtAuth := jwtProvider.New("secret")

	userRepository := memory.NewUserRepository()
	userService := services.NewUserService(userRepository)

	authService := services.NewAuthService(userRepository, jwtAuth)

	// route handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("success")
	})

	app.Post("/login", authHandler.Login)

	protected := app.Group("protected", middlewares.JwtAuth(jwtAuth))

	protected.Get("/update", middlewares.Updater(), func(c *fiber.Ctx) error {
		return c.SendString("protected for update")
	})

	protected.Get("/create", middlewares.Creator(), func(c *fiber.Ctx) error {
		return c.SendString("protected for create")
	})

	protected.Get("/delete", middlewares.Deletor(), func(c *fiber.Ctx) error {
		return c.SendString("protected for delete")
	})

	protected.Get("/read", middlewares.Reader(), func(c *fiber.Ctx) error {
		return c.SendString("protected for read")
	})

	userRouter := app.Group("users")
	userRouter.Post("/create", userHandler.CreateUser)
	userRouter.Get("/:id", userHandler.FindUserById)

	app.Listen(":9094")
}
