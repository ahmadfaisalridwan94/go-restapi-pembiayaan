package middleware

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"

	"pembiayaan/src/helpers"
)

type ErrorResponse struct {
	ResponseCode string `json:"ResponseCode"`
	ResponseDesc string `json:"ResponseDesc"`
}

// Authentication is a middleware function that handles authentication for HTTP requests.
// It validates the token, API key, and signature of the request.
// If the token, API key, or signature is invalid, it returns an error response.
// Otherwise, it sets the user and requestBody in the context and calls the next handler.
func Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)
			secretKey := os.Getenv("ACCESS_TOKEN_SECRET_KEY")

			if token == "" {
				helpers.ErrorMessage("0103", errors.New("invalid token"))
			}

			//validate jwt
			claims, err := helpers.ValidateJWT(&helpers.ParamsValidateJWT{
				Token:     token,
				SecretKey: secretKey,
			})

			if err != nil {
				return helpers.ErrorMessage("0103", err)
			}

			// Access the request body
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return helpers.ErrorMessage("0013", err)
			}

			// Seek back to the beginning of the request body
			c.Request().Body = io.NopCloser(bytes.NewBuffer(body))
			requestBody := string(body)

			apiKey := claims["apiKey"].(string)

			// get client secret
			clientSecret, err := helpers.DecryptApiKey(apiKey)
			if err != nil {
				return helpers.ErrorMessage("0104", errors.New("invalid api key"))
			}

			// check auth signature
			checkSignature := helpers.CheckAuthSignature(&helpers.AuthComponent{
				HttpMethod:   c.Request().Method,
				RequestPath:  c.Request().URL.String(),
				AccessToken:  token,
				ClientSecret: clientSecret,
				RequestBody:  requestBody,
				Signature:    c.Request().Header.Get("X-Signature"),
				Timestamp:    c.Request().Header.Get("X-Timestamp"),
			})

			if !checkSignature {
				return helpers.ErrorMessage("0101", errors.New("invalid signature"))
			}

			user := make(map[string]interface{})
			mapstructure.Decode(claims, &user)

			c.Set("user", user)
			c.Set("requestBody", requestBody)

			return next(c)
		}
	}
}
