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

func NewAuthController(srv *service.UserService) *AuthController {
	return &AuthController{srv}
}
