package routes

import (
	"github.com/Kei-K23/user-management-system-api/controllers"
	"github.com/Kei-K23/user-management-system-api/repositories"
	"github.com/Kei-K23/user-management-system-api/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	roleRepo := repositories.NewRoleRepository()
	roleService := services.NewRoleService(roleRepo)
	roleController := controllers.NewRoleController(*roleService)
	// Create Controller

	apiV1.Post("/roles", roleController.CreateRole)
}
