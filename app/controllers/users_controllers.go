package controllers

import (
	"mobileapp_go_fiber/app/models"
	"mobileapp_go_fiber/app/queries"
	"mobileapp_go_fiber/app/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetBalanceUserHandler(c *fiber.Ctx) error {
	now := time.Now().Unix()

	token, err := utils.VerifyToken(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	users := models.User{}
	users.JWTToken = token.Raw

	tokenUser := queries.IsJwtValid(users)
	if !tokenUser {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, invalid token",
		})
	}

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"login":   users.Phonenumber,
		"balance": users.Balance,
	})
}

func CreateUserHandler(c *fiber.Ctx) error {
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user := models.User{}
	user.CreatedAt = time.Now()
	user.JWTToken = token
	user.FName = c.FormValue("fname")
	user.Phonenumber = c.FormValue("phonenumber")
	user.Password = c.FormValue("password")
	user.Email = c.FormValue("email")
	user.Gender = c.FormValue("gender")
	user.Birthday = c.FormValue("birthday")
	user.Balance = 1000

	if err := queries.InsertUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}
