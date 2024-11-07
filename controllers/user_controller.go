package controllers

import (
	"strconv"

	"github.com/Kei-K23/user-management-system-api/dto"
	"github.com/Kei-K23/user-management-system-api/models"
	"github.com/Kei-K23/user-management-system-api/repositories"
	"github.com/Kei-K23/user-management-system-api/services"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (r *UserController) CreateUser(c *fiber.Ctx) error {
	var input dto.CreateUserInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := r.userService.Create(input.Username, input.FullName, input.Email, input.Password, input.RoleId)

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (r *UserController) GetUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when converting string to int"})
	}

	user, err := r.userService.GetById(id)

	if err == repositories.ErrUserNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when getting user"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (r *UserController) GetUsers(c *fiber.Ctx) error {
	var users []*models.User
	username := c.Query("username")

	if username != "" {
		user, err := r.userService.GetByUsername(username)

		if err == repositories.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when getting user by username"})
		}

		users = append(users, user)
	} else {
		usersData, err := r.userService.GetUsers()
		if err == repositories.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON([]models.User{})
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when getting users"})
		}
		users = append(users, usersData...)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (r *UserController) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when converting string to int"})
	}

	var input dto.UpdateUserInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	role, err := r.userService.Update(id, input.Username, input.FullName, input.Email, input.Password, input.RoleId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when updating user"})
	}

	return c.Status(fiber.StatusOK).JSON(role)
}

func (r *UserController) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when converting string to int"})
	}

	roleId, err := r.userService.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error when deleting user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": roleId})
}
