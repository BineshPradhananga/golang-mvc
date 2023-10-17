package middleware

import (
	"fmt"
	"time"

	"github.com/binesh/gomvc/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const SecretKey = "secret"

func JWTValidation() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			var response = helpers.UnauthorizedMessage("Missing or invalid token!!")
			statusCode := response.(map[string]interface{})["statusCode"]
			return c.Status(statusCode.(int)).JSON(response)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid token signing method")
			}
			return []byte(SecretKey), nil
		})

		if err != nil {
			var response = helpers.UnauthorizedMessage("Token is Invalid!!")
			statusCode := response.(map[string]interface{})["statusCode"]
			return c.Status(statusCode.(int)).JSON(response)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			// Check if the token has expired
			expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
			currentTime := time.Now()

			if currentTime.After(expirationTime) {
				var response = helpers.UnauthorizedMessage("Token has expired!!")
				statusCode := response.(map[string]interface{})["statusCode"]
				return c.Status(statusCode.(int)).JSON(response)
			}
			c.Locals("user", claims) // Store the claims in locals for access in route handlers
			return c.Next()
		}

		var response = helpers.UnauthorizedMessage("Unauthorized Access!!")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)
	}
}
