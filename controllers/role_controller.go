package controllers

import (
	"strconv"

	"github.com/Kei-K23/user-management-system-api/dto"
	"github.com/Kei-K23/user-management-system-api/services"
	"github.com/gofiber/fiber/v2"
)

type RoleController struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{roleService}
}

func (r *RoleController) CreateRole(c *fiber.Ctx) error {
	var input dto.CreateRoleInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	role, err := r.roleService.Create(input.Name, input.Description)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create role"})
	}

	return c.Status(fiber.StatusCreated).JSON(role)
}

func (r *RoleController) GetRoleById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when converting string to int"})
	}

	role, err := r.roleService.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when getting role"})
	}

	return c.Status(fiber.StatusOK).JSON(role)
}

func (r *RoleController) GetRoleByName(c *fiber.Ctx) error {
	name := c.Query("name")

	role, err := r.roleService.GetByName(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when getting role"})
	}

	return c.Status(fiber.StatusOK).JSON(role)
}

func (r *RoleController) UpdateRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when converting string to int"})
	}

	var input dto.UpdateRoleInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	role, err := r.roleService.Update(id, input.Name, input.Description)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when updating role"})
	}

	return c.Status(fiber.StatusOK).JSON(role)
}
