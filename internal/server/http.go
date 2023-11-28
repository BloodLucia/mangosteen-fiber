package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kalougata/bookkeeping/internal/controller"
	"github.com/kalougata/bookkeeping/pkg/config"
	"github.com/kalougata/bookkeeping/pkg/middleware"
)

func NewHTTPServer(
	authC *controller.AuthController,
	tagC *controller.TagController,
	itemC *controller.ItemController,
	jm *middleware.JWTMiddleware,
	conf *config.Config,
) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:       conf.APP.Name,
		StrictRouting: true,
		CaseSensitive: true,
	})

	v1Group := app.Group("/api/v1").Use(cors.New())

	noAuthGroup := v1Group.Group("")
	{
		noAuthGroup.Post("/login", authC.SignInWithEmail())
		noAuthGroup.Post("/sendVerificationCode", authC.SendVerificationCode())
	}

	authGroup := v1Group.Group("").Use(jm.JWTAuth())
	{
		authGroup.Get("/ping", authC.Ping())
		authGroup.Get("/me", authC.Me())

		// 标签路由
		tagsGroup := authGroup.Group("/tags")
		{
			tagsGroup.Get("", tagC.List())
			tagsGroup.Post("/create", tagC.Create())
		}

		// 账单路由
		itemsGroup := authGroup.Group("/items")
		{
			itemsGroup.Get("", itemC.List())
			itemsGroup.Post("/create", itemC.Create())
		}

		authGroup.Get("/balance", itemC.Balance())
	}

	return app
}
