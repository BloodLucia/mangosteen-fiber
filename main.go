package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to listen Serve err: %s \n", err)
	}
}
