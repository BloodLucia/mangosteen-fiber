package model

import (
	"github.com/gookit/goutil"
	"time"
)

type Item struct {
	ID         uint64    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt  time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt  time.Time `xorm:"updated TIMESTAMP updated_at"`
	Amount     int
	Type       string
	TagId      uint64
	UserId     uint64
	HappenedAt string
}

func (i *Item) TableName() string {
	return "t_items"
}

type ItemInReq struct {
	Amount int    `json:"amount" validate:"int"`
	Type   string `json:"type"`
	TagId  string `json:"tagId"`
	UserId string `json:"userId"`
}

type ItemListReq struct {
	HappenedAfter  string `query:"happened_after"`
	HappenedBefore string `query:"happened_before"`
	Page           int    `query:"page"`
	Limit          int    `query:"limit"`
	UserId         uint64 `query:"user_id"`
}

func (u *ItemInReq) ToModel() *Item {
	return &Item{
		Amount: u.Amount,
		TagId:  goutil.Uint(u.TagId),
		UserId: goutil.Uint(u.UserId),
	}
}
