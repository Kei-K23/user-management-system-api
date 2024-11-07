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

	// User repository setup
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(*userService)
	apiV1.Get("/users", userController.GetUsers)
	apiV1.Get("/users/:id", userController.GetUserById)
	apiV1.Post("/users", userController.CreateUser)
	apiV1.Put("/users/:id", userController.UpdateUser)
	apiV1.Delete("/users/:id", userController.DeleteUser)

	// Auth
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(*authService)
	apiV1.Post("/auth/login", authController.Login)
	apiV1.Post("/auth/register", authController.Register)
}
