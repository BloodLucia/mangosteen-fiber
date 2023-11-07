package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/goutil"
	"github.com/kalougata/bookkeeping/internal/dto"
	"github.com/kalougata/bookkeeping/internal/model"
	"github.com/kalougata/bookkeeping/internal/service"
	"github.com/kalougata/bookkeeping/pkg/e"
	"github.com/kalougata/bookkeeping/pkg/response"
	"github.com/kalougata/bookkeeping/pkg/validator"
	"net/http"
)

type ItemController struct {
	service *service.ItemService
}

func (ic *ItemController) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := &model.ItemInReq{UserId: ctx.GetRespHeader("userId")}
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

func (ic *ItemController) List() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		query := &model.ItemListReq{UserId: goutil.Uint(ctx.GetRespHeader("userId"))}
		if err := ctx.QueryParser(query); err != nil {
			return response.Handle(ctx, e.ErrBadRequest().WithErr(err), nil)
		}
		list, err := ic.service.List(ctx.Context(), query)
		if err != nil {
			return response.Handle(ctx, e.ErrInternalServer().WithErr(err), nil)
		}

		return response.Handle(ctx, nil, list)
	}
}

func (ic *ItemController) Balance() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		queries := &dto.BalanceQueries{}
		if err := ctx.QueryParser(queries); err != nil {
			return response.Handle(ctx, e.New(http.StatusUnprocessableEntity, err.Error()), nil)
		}
		resp, err := ic.service.Balance(ctx.Context(), queries)
		if err != nil {
			return response.Handle(ctx, err, nil)
		}

		return response.Handle(ctx, nil, resp)
	}
}

func NewItemController(service *service.ItemService) *ItemController {
	return &ItemController{service}
}
