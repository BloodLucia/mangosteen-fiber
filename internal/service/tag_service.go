package service

import (
	"context"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/dto"
	"github.com/kalougata/bookkeeping/internal/model"
	"github.com/kalougata/bookkeeping/pkg/e"
)

type TagService struct {
	data *data.Data
}

func (ts *TagService) Create(ctx context.Context, req *dto.TagInBody) error {
	if exist, err := ts.data.DB.Context(ctx).Table(&model.Tag{}).Where("name = ? AND type = ?", req.Name, req.Type).Exist(); exist && err == nil {
		return e.ErrBadRequest().WithMsg("标签名已存在~")
	}
	if count, err := ts.data.DB.Context(ctx).Table(&model.Tag{}).Insert(req.ToModel()); err != nil || count <= 0 {
		return e.ErrInternalServer().WithMsg("添加标签失败, 请稍后再试~").WithErr(err)
	}

	return nil
}

func (ts *TagService) List(ctx context.Context, userId string) (list []*model.Tag, err error) {
	list = make([]*model.Tag, 0)
	err = ts.data.DB.Context(ctx).Table(&model.Tag{}).Where("user_id = ?", userId).Find(&list)

	return
}

func NewTagService(data *data.Data) *TagService {
	return &TagService{data}
}
