package routes

import (
	"github.com/Kei-K23/user-management-system-api/controllers"
	"github.com/Kei-K23/user-management-system-api/repositories"
	"github.com/Kei-K23/user-management-system-api/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	// Role repository setup
	roleRepo := repositories.NewRoleRepository()
	roleService := services.NewRoleService(roleRepo)
	roleController := controllers.NewRoleController(*roleService)
	apiV1.Get("/roles", roleController.GetRoles)
	apiV1.Get("/roles/:id", roleController.GetRoleById)
	apiV1.Post("/roles", roleController.CreateRole)
	apiV1.Put("/roles/:id", roleController.UpdateRole)
	apiV1.Delete("/roles/:id", roleController.DeleteRole)
}
