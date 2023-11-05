package service

import (
	"context"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/model"
	"github.com/kalougata/bookkeeping/pkg/e"
)

type UserService struct {
	data *data.Data
}

func (us *UserService) FindOrCreateWithEmail(ctx context.Context, req *model.UserInReq) error {
	if req.VerificationCode != "123456" {
		return e.ErrBadRequest().WithMsg("验证码错误")
	}
	exist, err := us.data.DB.Context(ctx).Table(&model.User{}).Where("email = ?", req).Exist()
	if err != nil {
		return e.ErrInternalServer().WithErr(err)
	}
	if !exist {
		// 创建用户
		if count, err := us.data.DB.Context(ctx).Table(&model.User{}).Insert(req.ToModel()); err != nil || count <= 0 {
			return e.ErrInternalServer().WithErr(err)
		}
	}

	// 1. 校验密码
	return nil
}

func NewUserService(data *data.Data) *UserService {
	return &UserService{data}
}
