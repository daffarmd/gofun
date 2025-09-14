package handler

import (
	"github.com/daffarmd/gofun/backend-fun/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	services *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	h.services.Register(req.Username, req.Password)
	return c.JSON(fiber.Map{"message": "Register Success"})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if h.services.Login(req.Username, req.Password) {
		return c.Status(200).JSON(fiber.Map{"message": "Login Success"})
	}

	return c.Status(401).JSON(fiber.Map{"error": "Invalid Credential"})
}
