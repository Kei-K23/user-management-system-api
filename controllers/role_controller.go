package controllers

import (
	"github.com/Kei-K23/user-management-system-api/models"
	"github.com/Kei-K23/user-management-system-api/services"
	"github.com/gofiber/fiber/v2"
)

func CreateRole(roleService services.RoleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get and parse request body
		var req map[string]string
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		role := models.Role{
			Name:        req["name"],
			Description: req["description"],
		}

		if err := roleService.CreateRole(role); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to create role")
		}
		return c.Status(fiber.StatusCreated).JSON(role)
	}
}
