package middlewares

import (
	"client_administration/services/jwt"

	"github.com/gofiber/fiber"
)


func IsLogedIn(c *fiber.Ctx){

    authHeader := c.Get("Authorization")
	if authHeader == "" {
		 c.Status(fiber.StatusUnauthorized).Send((fiber.Map{
			"message": "Missing Authorization header",
		}))
		return
	}
	_, userInfor, err := jwt.VeryfiToken(authHeader)

	if err != nil {
		c.Status(fiber.StatusUnauthorized).Send((fiber.Map{
			"message": "Invalid or expired token",
		}))
		return
	}
	
	c.Locals("user", userInfor)
	c.Next()
}