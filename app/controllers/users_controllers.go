package controllers

import (
	"math/rand"
	"mobileapp_go_fiber/app/models"
	"mobileapp_go_fiber/app/queries"
	"mobileapp_go_fiber/app/utils"
	"strconv"
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

	user, err := queries.FindUserToken(token.Raw)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, invalid token",
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"login":   user.Phonenumber,
		"balance": user.Balance,
	})
}

func GetUserTransactionsHandler(c *fiber.Ctx) error {
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

	user, err := queries.FindUserToken(token.Raw)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, invalid token",
		})
	}

	trcs, err := queries.GetUserTransactions(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"transactions": trcs,
	})
}

func PayHandler(c *fiber.Ctx) error {
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

	user, err := queries.FindUserToken(token.Raw)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, invalid token",
		})
	}

	trcs := models.Transaction{}
	trcs.ID = rand.Intn(100000)
	trcs.IDUser = user.ID
	trcs.CreatedAt = time.Now()
	trcs.Phonenumber = c.Params("phonenumber")
	trcs.Summary, _ = strconv.Atoi(c.Params("sum"))

	err = queries.InsertTransaction(trcs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if trcs.Summary > user.Balance {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "summary transaction cannot be bigger balance",
		})
	} else if user.Balance < 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "balance cannot be minus",
		})
	}

	user.Balance -= trcs.Summary

	err = queries.UpdateBalance(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "payment succesful",
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
	user.ID = rand.Intn(100000)
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

func UpdateUserDataHandler(c *fiber.Ctx) error {
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

	user, err := queries.FindUserToken(token.Raw)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, invalid token",
		})
	}

	user.FName = c.FormValue("fname")
	user.Email = c.FormValue("email")
	user.Gender = c.FormValue("gender")
	user.Birthday = c.FormValue("birthday")

	err = queries.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "user dates is updates",
	})
}
