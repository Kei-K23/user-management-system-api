package controllers

import (
	"log"

	"github.com/Kei-K23/user-management-system-api/dto"
	"github.com/Kei-K23/user-management-system-api/services"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

func (a *AuthController) Register(c *fiber.Ctx) error {
	var input dto.CreateUserInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := a.authService.Register(input.Username, input.FullName, input.Email, input.Password, input.RoleId)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (a *AuthController) Login(c *fiber.Ctx) error {
	var input dto.LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	token, err := a.authService.Login(input.Username, input.Password)

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid auth credentials"})
	}

	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to login"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"access_token": token})
}
