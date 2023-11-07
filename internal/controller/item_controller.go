package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/goutil"
	"github.com/kalougata/bookkeeping/internal/model"
	"github.com/kalougata/bookkeeping/internal/service"
	"github.com/kalougata/bookkeeping/pkg/e"
	"github.com/kalougata/bookkeeping/pkg/response"
	"github.com/kalougata/bookkeeping/pkg/validator"
)

type ItemController struct {
	service *service.ItemService
}

func (ic *ItemController) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := &model.ItemInReq{}
		if err := validator.Checker(ctx, data); err != nil {
			return response.Handle(ctx, err, nil)
		}
		if !goutil.IsEqual(data.UserId, ctx.GetRespHeader("userId")) {
			return response.Handle(ctx, e.ErrForbidden(), nil)
		}
		if c, err := ic.service.Create(ctx.Context(), data); err != nil || c <= 0 {
			return response.Handle(ctx, e.ErrInternalServer().WithErr(err), nil)
		}

		return response.Handle(ctx, nil, data)
	}
}

func NewItemController(service *service.ItemService) *ItemController {
	return &ItemController{service}
}
