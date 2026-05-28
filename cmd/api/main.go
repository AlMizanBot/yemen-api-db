package main

import (
    "github.com/gofiber/fiber/v2"
    "my-gateway/internal/auth"
    "my-gateway/pkg/config"
)

func main() {
    cfg := config.LoadConfig()
    app := fiber.New()

    api := app.Group("/api/v1")
    api.Use(auth.APIKeyAuth) // الحماية

    api.Post("/topup", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"success": true, "message": "تمت العملية"})
    })

    app.Listen(":" + cfg.Port)
}
