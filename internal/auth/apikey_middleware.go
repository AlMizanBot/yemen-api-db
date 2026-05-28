package auth

import "github.com/gofiber/fiber/v2"

func APIKeyAuth(c *fiber.Ctx) error {
    apiKey := c.Get("X-API-KEY")
    if apiKey == "" {
        return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
    }
    return c.Next()
}
