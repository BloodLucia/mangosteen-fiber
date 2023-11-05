package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/bookkeeping/internal/controller"
)

func NewHTTPServer(
	authC *controller.AuthController,
	tagC *controller.TagController,
	itemC *controller.ItemController,
) *fiber.App {
	app := fiber.New()

	app.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString("HelloWorld")
	})

	return app
}
