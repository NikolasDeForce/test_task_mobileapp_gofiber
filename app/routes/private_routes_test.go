package routes

import (
	"io"
	"mobileapp_go_fiber/app/utils"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPrivateRoutes(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		panic(err)
	}

	tests := []struct {
		description   string
		route         string
		method        string
		tokenString   string
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "GET balance without JWT",
			route:         "/api/v1/balance",
			method:        "GET",
			tokenString:   "",
			body:          nil,
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "GET balance with JWT",
			route:         "/api/v1/balance",
			method:        "GET",
			tokenString:   "Bearer " + token,
			body:          nil,
			expectedError: false,
			expectedCode:  401, //unauthorized
		},
		{
			description:   "GET history with JWT",
			route:         "/api/v1/history",
			method:        "GET",
			tokenString:   "Bearer " + token,
			body:          nil,
			expectedError: false,
			expectedCode:  401, //unauthorized
		},
		{
			description:   "POST pay with JWT",
			route:         "/api/v1/pay/79999999/50",
			method:        "POST",
			tokenString:   "Bearer " + token,
			body:          nil,
			expectedError: false,
			expectedCode:  401, //unauthorized
		},
	}

	app := fiber.New()

	PrivateRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, nil)
		req.Header.Set("Authorization", test.tokenString)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
