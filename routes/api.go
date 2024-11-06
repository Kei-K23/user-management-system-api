package routes

import (
	"github.com/Kei-K23/user-management-system-api/controllers"
	"github.com/Kei-K23/user-management-system-api/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, roleService services.RoleService) {
	// Role CRUD
	app.Post("/roles", controllers.CreateRole(roleService))
}
