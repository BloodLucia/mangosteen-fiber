package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/bookkeeping/internal/model"
	"github.com/kalougata/bookkeeping/internal/service"
	"github.com/kalougata/bookkeeping/pkg/response"
	"github.com/kalougata/bookkeeping/pkg/validator"
)

type AuthController struct {
	srv *service.UserService
}

func (ac *AuthController) SignInWithEmail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := &model.UserInReq{}
		if err := validator.Checker(ctx, data); err != nil {
			return response.Handle(ctx, err, nil)
		}
		if resp, err := ac.srv.FindOrCreateWithEmail(ctx.Context(), data); err != nil {
			return response.Handle(ctx, err, nil)
		} else {
			return response.Handle(ctx, nil, resp)
		}
	}
}

func (ac *AuthController) SendVerificationCode() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := &model.UserSendEmailReq{}
		if err := validator.Checker(ctx, data); err != nil {
			return response.Handle(ctx, err, nil)
		}
		if err := ac.srv.SendVerificationCode(ctx.Context(), data); err != nil {
			return response.Handle(ctx, err, nil)
		}

		return response.Handle(ctx, nil, "请查看邮箱")
	}
}

func (ac *AuthController) Ping() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return response.Handle(ctx, nil, ctx.GetRespHeader("userId"))
	}
}

func NewAuthController(srv *service.UserService) *AuthController {
	return &AuthController{srv}
}
