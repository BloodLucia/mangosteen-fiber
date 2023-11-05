package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/pkg/response"
)

type AuthController struct {
	data *data.Data
}

func (ac *AuthController) SignInWithEmail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return response.Handle(ctx, nil, "hello!")
	}
}

func NewAuthController(data *data.Data) *AuthController {
	return &AuthController{data}
}
