package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuth(context fiber.Ctx) error {
	// get cookie 
	tokenString := context.Cookies("Authorization")

	// if cookie not exist make unauthorized request
	if tokenString == "" {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "resource access forbidden"})
	}

	// parse token and check validation
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil || !token.Valid{
		return context.Status(http.StatusUnauthorized).
        JSON(fiber.Map{"message": "invalid or expired token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims);

	if !ok {
		return context.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "failed to claimed token"})
	}

	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return context.Status(http.StatusForbidden).JSON(fiber.Map{"message": "token expired, please login"})
	}

	context.Locals("user", claims["sig"])

	return context.Next()
}