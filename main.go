package main

import (
	"mobileapp_go_fiber/app/configs"
	"mobileapp_go_fiber/app/middleware"
	"mobileapp_go_fiber/app/routes"
	"mobileapp_go_fiber/app/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	utils.StartServerWithGracefulShutdown(app)
}
