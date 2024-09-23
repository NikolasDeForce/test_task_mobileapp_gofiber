package routes

import (
	"mobileapp_go_fiber/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/register", controllers.CreateUserHandler)
}
