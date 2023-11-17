package model

import (
	"time"
)

type Tag struct {
	ID        uint64    `xorm:"not null pk autoincr BIGINT(20) id" json:"id"`
	Name      string    `xorm:"not null VARCHAR(30) name" json:"name"`
	Kind      string    `xorm:"not null VARCHAR(10) kind" json:"kind"`
	Sign      string    `xorm:"not null CHAR(1) sign" json:"sign"`
	UserId    uint64    `xorm:"not null BIGINT(20) user_id" json:"user_id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at" json:"updated_at"`
}

func (t *Tag) TableName() string {
	return "t_tags"
}
