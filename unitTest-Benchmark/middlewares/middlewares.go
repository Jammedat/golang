package middleware

import "github.com/gofiber/fiber/v2"


func SecurityHeaders(c *fiber.Ctx) error {
	c.Response().Header.Add("Content-Security-Policy", "default-src 'self'")
	c.Response().Header.Add("X-Frame-Options", "Deny")
	return c.Next()
	
}

