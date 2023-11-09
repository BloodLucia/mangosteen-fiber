package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/goutil"
	"github.com/kalougata/bookkeeping/internal/dto"
	"github.com/kalougata/bookkeeping/internal/service"
	"github.com/kalougata/bookkeeping/pkg/e"
	"github.com/kalougata/bookkeeping/pkg/response"
	"github.com/kalougata/bookkeeping/pkg/validator"
	"net/http"
)

type TagController struct {
	service *service.TagService
}

func (tc *TagController) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := &dto.TagInBody{}
		userId := ctx.GetRespHeader("userId")
		if goutil.IsEmpty(userId) {
			return response.Handle(ctx, e.ErrUnauthorized(), nil)
		}
		if err := validator.Checker(ctx, data); err != nil {
			return response.Handle(ctx, err, nil)
		}
		if err := tc.service.Create(ctx.Context(), data); err != nil {
			return response.Handle(ctx, err, nil)
		}

		return response.Handle(ctx, nil, data)
	}
}

func (tc *TagController) List() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Query("userId")
		queries := &dto.TagListQueries{}
		if err := ctx.QueryParser(queries); err != nil {
			return response.Handle(ctx, e.New(http.StatusUnprocessableEntity, err.Error()), nil)
		}
		if !goutil.IsEqual(userId, ctx.GetRespHeader("userId")) {
			return response.Handle(ctx, e.ErrForbidden(), nil)
		}
		list, err := tc.service.List(ctx.Context(), queries)
		if err != nil {
			return response.Handle(ctx, e.ErrInternalServer().WithErr(err), nil)
		}

		return response.Handle(ctx, nil, list)
	}
}

func NewTagController(srv *service.TagService) *TagController {
	return &TagController{srv}
}
