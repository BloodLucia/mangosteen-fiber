package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/bookkeeping/internal/service"
	"github.com/kalougata/bookkeeping/pkg/response"
)

type AuthController struct {
	srv *service.UserService
}

func (ac *AuthController) SignInWithEmail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return response.Handle(ctx, nil, "hello!")
	}
}

func NewAuthController(srv *service.UserService) *AuthController {
	return &AuthController{srv}
}
