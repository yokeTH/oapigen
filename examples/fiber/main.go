package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()

	app.Get("healthz", func(ctx fiber.Ctx) error {
		return ctx.SendString("OK")
	})
}
