package model

import "time"

type User struct {
	ID        uint64    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt time.Time `xorm:"deleted DATETIME deleted_at"`
	Email     string    `xorm:"not null VARCHAR(100) unique index email"`
}

type UserInReq struct {
	VerificationCode string `json:"verificationCode" validate:"required|minLen:6|maxLen:6" message:"required:{field} 必填|minLen:{field} 验证码长度是6个字符|maxLen:{field} 验证码长度是6个字符"`
	Email            string `json:"email" validate:"required|email" message:"required:{field} 必填|email:{field} 邮箱格式错误"`
}

type UserOutRes struct {
	UserId uint64 `json:"userId"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

func (u User) TableName() string {
	return "t_users"
}

func (r *UserInReq) ToModel() *User {
	return &User{
		Email: r.Email,
	}
}
