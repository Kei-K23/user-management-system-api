package main

import (
	"github.com/Kei-K23/user-management-system-api/config"
	"github.com/Kei-K23/user-management-system-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	godotenv.Load()

	// Connect to db
	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3030")
}
