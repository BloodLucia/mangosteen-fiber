package model

import "time"

type User struct {
	ID        uint64    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt time.Time `xorm:"deleted DATETIME deleted_at"`
	Email     string
}

type UserInReq struct {
	VerificationCode string `json:"verificationCode"`
	Email            string `json:"email"`
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
