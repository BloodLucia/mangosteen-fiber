package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gookit/goutil"
	"github.com/kalougata/bookkeeping/internal/data"
	"github.com/kalougata/bookkeeping/internal/dto"
	"github.com/kalougata/bookkeeping/internal/model"
	"github.com/kalougata/bookkeeping/pkg/e"
	"github.com/kalougata/bookkeeping/pkg/jwt"
	"github.com/kalougata/bookkeeping/pkg/mailer"
)

type UserService struct {
	data   *data.Data
	jwt    *jwt.JWT
	mailer *mailer.Mailer
}

func (us *UserService) SendVerificationCode(ctx context.Context, email string) (err error) {
	tmpCode := "123456"
	options := &mailer.Options{
		To:      email,
		Subject: "邮箱验证码",
		Text:    fmt.Sprintf("您的验证码为：%s，五分钟后过期。", tmpCode),
	}
	if err := us.mailer.Send(options); err != nil {
		return e.ErrInternalServer().WithMsg("发送验证码失败, 请稍后再试~").WithErr(err)
	}
	if err = us.data.Cache.Set(ctx, email, tmpCode, time.Minute*5).Err(); err != nil {
		return e.ErrInternalServer().WithErr(err)
	}
	return nil
}

func (us *UserService) FindOrCreate(ctx context.Context, req *dto.UserInBody) (*dto.UserOutBody, error) {
	user := &model.User{}
	//// 1. 从redis获取验证码
	val := us.data.Cache.Get(ctx, req.Email).Val()
	if goutil.IsEmpty(val) || !goutil.IsEqual(val, req.VerificationCode) {
		return nil, e.ErrBadRequest().WithMsg("验证码错误或已失效")
	}

	// 2. 查询用户
	exist, err := us.data.DB.Context(ctx).Table(&model.User{}).Where("email = ?", req.Email).Get(user)
	if err != nil {
		return nil, e.ErrInternalServer().WithErr(err)
	}

	claims := jwt.MyCustomClaims{UserId: goutil.String(user.ID)}
	token, _ := us.jwt.BuildToken(claims)
	resp := &dto.UserOutBody{
		UserId: user.ID,
		Email:  user.Email,
		Token:  token,
	}

	// 3. 用户不存在，创建用户
	if !exist {
		log.Println(fmt.Sprintf("用户不存在: %s，创建用户", req.Email))
		user = req.ToModel()
		if count, err := us.data.DB.Context(ctx).Table(&model.User{}).Insert(user); err != nil || count <= 0 {
			log.Println(err)
			return nil, e.ErrInternalServer().WithErr(err)
		}
		resp.UserId = user.ID
		resp.Email = user.Email

		return resp, nil
	}
	log.Println(fmt.Sprintf("用户已存在：%s，直接颁发token", user.Email))

	return resp, nil
}

func (us *UserService) GetCurrentUser(ctx context.Context, userId string) (*dto.UserProfileResp, error) {
	user := &model.User{}
	has, err := us.data.DB.Context(ctx).Table(&model.User{}).Where("id = ?", userId).Get(user)
	if err != nil {
		return nil, e.ErrInternalServer().WithErr(err)
	}
	if !has {
		return nil, e.ErrUnauthorized()
	}
	resp := &dto.UserProfileResp{
		UserId:    user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.GoString(),
	}

	return resp, nil
}

func NewUserService(data *data.Data, jwt *jwt.JWT, mailer *mailer.Mailer) *UserService {
	return &UserService{data, jwt, mailer}
}
