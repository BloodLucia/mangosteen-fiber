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
	// 1. 从redis获取验证码
	val := us.data.Cache.Get(ctx, req.Email).Val()
	if val == "" || val != req.VerificationCode {
		return e.ErrBadRequest().WithMsg("验证码错误或已失效")
	}

	// 2. 查询用户
	exist, err := us.data.DB.Context(ctx).Table(&model.User{}).Where("email = ?", req).Exist()
	if err != nil {
		return e.ErrInternalServer().WithErr(err)
	}
	// 3. 用户不存在，创建用户
	if !exist {
		if count, err := us.data.DB.Context(ctx).Table(&model.User{}).Insert(req.ToModel()); err != nil || count <= 0 {
			return e.ErrInternalServer().WithErr(err)
		}
	}

	return nil
}

func NewUserService(data *data.Data) *UserService {
	return &UserService{data}
}
