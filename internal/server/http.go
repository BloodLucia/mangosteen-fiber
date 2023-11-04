package server

import "github.com/gofiber/fiber/v2"

func NewHTTPServer() *fiber.App {
	app := fiber.New()

	app.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString("HelloWorld")
	})

	return app
}
