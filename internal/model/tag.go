package model

import (
	"time"
)

type Tag struct {
	ID        uint64    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	Name      string    `xorm:"not null VARCHAR(30) name"`
	Kind      string    `xorm:"not null VARCHAR(10) kind"`
	Sign      string    `xorm:"not null CHAR(1) sign"`
	UserId    uint64    `xorm:"not null BIGINT(20) user_id"`
}

func (u *Tag) TableName() string {
	return "t_tags"
}
