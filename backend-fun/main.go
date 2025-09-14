package main

import (
	"github.com/daffarmd/gofun/backend-fun/handler"
	"github.com/daffarmd/gofun/backend-fun/repository"
	"github.com/daffarmd/gofun/backend-fun/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Public routes
	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	// Protected routes
	// protected := app.Group("/api")
	// protected.Use(middleware.AuthMiddleware())

	app.Listen(":8081")
}
