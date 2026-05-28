package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"my-gateway/internal/auth"
	"my-gateway/pkg/config"
)

func main() {
	cfg := config.LoadConfig()

	// 1. الاتصال بقاعدة البيانات باستخدام Connection Pooling
	dbPool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbPool.Close()

	// التأكد من أن الاتصال فعال
	err = dbPool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}
	log.Println("Successfully connected to PostgreSQL on Render!")

	app := fiber.New()

	// 2. تمرير الـ dbPool للـ Middleware أو الـ Handlers لاحقاً
	// (مبدئياً نستخدمه للتحقق من المفاتيح)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("System is healthy and DB is connected")
	})

	api := app.Group("/api/v1")
	api.Use(auth.APIKeyAuth)

	api.Post("/topup", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ready to process transactions"})
	})

	app.Listen(":" + cfg.Port)
}
