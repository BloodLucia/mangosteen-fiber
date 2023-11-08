package service

import (
	"context"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/dto"
	"github.com/kalougata/bookkeeping/internal/model"
	"github.com/kalougata/bookkeeping/pkg/e"
	"github.com/kalougata/bookkeeping/pkg/page"
)

type ItemService struct {
	data *data.Data
}

func (is *ItemService) Create(ctx context.Context, req *dto.ItemInBody) (int64, error) {
	return is.data.DB.Context(ctx).Table(&model.Item{}).Insert(req.ToModel())
}

func (is *ItemService) List(ctx context.Context, req *dto.ItemListQueries) (*page.Page[*model.Item], error) {
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

func (is *ItemService) Balance(ctx context.Context, req *dto.BalanceQueries) (resp *dto.BalanceResponse, err error) {
	// 1. 获取所有账单
	list := make([]*model.Item, 0)
	err = is.data.DB.Context(ctx).Table(&model.Item{}).Where("user_id = ?", req.UserId).Find(&list)
	if err != nil {
		return nil, e.ErrInternalServer().WithErr(err)
	}

	var totalIncome float64
	var totalExpenses float64
	var totalBalance float64

	totalIncome, err = is.data.DB.Context(ctx).Table(&model.Item{}).Where("user_id = ? AND type = ?", 2, "income").Sum(&model.Item{}, "amount")
	totalExpenses, err = is.data.DB.Context(ctx).Table(&model.Item{}).Where("user_id = ? AND type = ?", 2, "expenses").Sum(&model.Item{}, "amount")
	if err != nil {
		return nil, e.ErrInternalServer().WithErr(err)
	}

	// 2. 计算总支出/总收入

	// 3. 计算净收入
	totalBalance = totalIncome - totalBalance

	resp = &dto.BalanceResponse{
		Income:   totalIncome,
		Expenses: totalExpenses,
		Balance:  totalBalance,
	}

	return resp, nil
}

func (is *ItemService) GetTotalAmountByIncome(ctx context.Context, userId uint64) (float64, error) {
	return is.data.DB.Context(ctx).Table(&model.Item{}).Where("user_id = ? AND type = ?", userId, "income").Sum(nil, "amount")
}

func (is *ItemService) GetTotalAmountByExpenses(ctx context.Context, userId uint64) (float64, error) {
	return is.data.DB.Context(ctx).Table(&model.Item{}).Where("user_id = ? AND type = ?", userId, "expenses").Sum(nil, "amount")
}

func NewItemService(data *data.Data) *ItemService {
	return &ItemService{data}
}
