package dto

import "github.com/kalougata/bookkeeping/internal/model"

type UserInBody struct {
	Email            string `json:"email" validate:"required|email" message:"required:{field} 必填|email:{field} 邮箱格式错误"`
	VerificationCode string `json:"verification_code" validate:"required|minLen:6|maxLen:6" message:"required:{field} 必填|minLen:{field} 验证码长度是6个字符|maxLen:{field} 验证码长度是6个字符"`
}

type UserEmailBody struct {
	Email string `json:"email" validate:"required|email" message:"required:{field} 必填|email:{field} 邮箱格式错误"`
}

type UserOutBody struct {
	UserId uint64 `json:"userId"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

func (r *UserInBody) ToModel() *model.User {
	return &model.User{
		Email: r.Email,
	}
}
