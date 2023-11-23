package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/bookkeeping/internal/dto"
	"github.com/kalougata/bookkeeping/internal/service"
	"github.com/kalougata/bookkeeping/pkg/response"
	"github.com/kalougata/bookkeeping/pkg/validator"
)

type AuthController struct {
	service *service.UserService
}

func (ac *AuthController) SignInWithEmail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := &dto.UserInBody{}
		if err := validator.Checker(ctx, data); err != nil {
			return response.Handle(ctx, err, nil)
		}
		resp, err := ac.service.FindOrCreate(ctx.Context(), data)
		if err != nil {
			return response.Handle(ctx, err, nil)
		}

		return response.Handle(ctx, nil, resp)
	}
}

func (ac *AuthController) SendVerificationCode() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := &dto.UserEmailBody{}
		if err := validator.Checker(ctx, data); err != nil {
			return response.Handle(ctx, err, nil)
		}
		if err := ac.service.SendVerificationCode(ctx.Context(), data); err != nil {
			return response.Handle(ctx, err, nil)
		}

		return response.Handle(ctx, nil, nil)
	}
}

func (ac *AuthController) Ping() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return response.Handle(ctx, nil, fmt.Sprintf("用户ID: %v", ctx.GetRespHeader("userId")))
	}
}

func NewAuthController(srv *service.UserService) *AuthController {
	return &AuthController{srv}
}
