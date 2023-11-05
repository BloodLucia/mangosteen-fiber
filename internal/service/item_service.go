package service

import (
	"context"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/model"
)

type ItemService struct {
	data *data.Data
}

func (is *ItemService) Create(ctx context.Context, req *model.ItemInReq) (int64, error) {
	return is.data.DB.Context(ctx).Table(&model.Item{}).Insert(req.ToModel())
}

func NewItemService(data *data.Data) *ItemService {
	return &ItemService{data}
}
