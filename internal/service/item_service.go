package service

import "github.com/kalougata/bookkeeping/internal/data"

type ItemService struct {
	data *data.Data
}

func NewItemService(data *data.Data) *ItemService {
	return &ItemService{data}
}
