package controller

import "github.com/kalougata/bookkeeping/internal/service"

type ItemController struct {
	service *service.ItemService
}

func NewItemController(service *service.ItemService) *ItemController {
	return &ItemController{service}
}
