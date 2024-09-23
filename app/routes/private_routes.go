package routes

import (
	"mobileapp_go_fiber/app/controllers"
	"mobileapp_go_fiber/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjcxNzEyMzZ9.DwPOmW0EfaNiib1dhndlYndQ9MxkFfJRw6h4hCEh7ho

	route.Get("/balance", middleware.JWTProtected(), controllers.GetBalanceUserHandler)
	route.Get("/history", middleware.JWTProtected(), controllers.GetUserTransactionsHandler)

	route.Put("/profile/update", middleware.JWTProtected(), controllers.UpdateUserDataHandler)

	route.Post("/pay/:phonenumber/:sum", middleware.JWTProtected(), controllers.PayHandler)
}
