package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/bookkeeping/internal/service"
	"github.com/kalougata/bookkeeping/pkg/response"
)

type ItemController struct {
	service *service.ItemService
}

func (ic *ItemController) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return response.Handle(ctx, nil, "create item")
	}
}

func NewItemController(service *service.ItemService) *ItemController {
	return &ItemController{service}
}
