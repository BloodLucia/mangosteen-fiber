package service

import (
	"context"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/model"
	"github.com/kalougata/bookkeeping/pkg/e"
	"github.com/kalougata/bookkeeping/pkg/page"
)

type ItemService struct {
	data *data.Data
}

func (is *ItemService) Create(ctx context.Context, req *model.ItemInReq) (int64, error) {
	return is.data.DB.Context(ctx).Table(&model.Item{}).Insert(req.ToModel())
}

func (is *ItemService) List(ctx context.Context, req *model.ItemListReq) (*page.Page[*model.Item], error) {
	list := make([]*model.Item, 0)
	err := is.data.DB.Context(ctx).Table(&model.Item{}).Where("user_id = ? AND happened_at >= ? AND happened_at <= ?", req.UserId, req.HappenedAfter, req.HappenedBefore).Limit(10, req.Page).Find(&list)
	if err != nil {
		return nil, e.ErrInternalServer().WithErr(err)
	}
	p := page.Build[*model.Item](list, &page.Pager{
		Page:    req.Page,
		PerPage: 10,
		Count:   len(list),
	})

	return p, nil
}

func NewItemService(data *data.Data) *ItemService {
	return &ItemService{data}
}
