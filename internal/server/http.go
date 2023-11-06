package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/bookkeeping/internal/controller"
	"github.com/kalougata/bookkeeping/pkg/middleware"
)

func NewHTTPServer(
	authC *controller.AuthController,
	tagC *controller.TagController,
	itemC *controller.ItemController,
	jm *middleware.JWTMiddleware,
) *fiber.App {
	app := fiber.New()

	v1Group := app.Group("/api/v1")

	noAuthGroup := v1Group.Group("")
	{
		noAuthGroup.Post("/login", authC.SignInWithEmail())

		noAuthGroup.Post("/sendVerificationCode", authC.SendVerificationCode())
	}

	authGroup := v1Group.Group("").Use(jm.JWTAuth())
	{
		authGroup.Get("/ping", func(ctx *fiber.Ctx) error {
			return ctx.SendString("pong!")
		})
	}

	return app
}
