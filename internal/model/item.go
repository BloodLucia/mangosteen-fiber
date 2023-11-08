package model

import (
	"time"
)

type Item struct {
	ID         uint64    `json:"id" xorm:"not null pk autoincr BIGINT(20) id"`
	Amount     int       `json:"amount" xorm:"not null DECIMAL(10,2) amount"`
	Kind       string    `json:"type" xorm:"not null VARCHAR(10) kind"`
	TagId      uint64    `json:"tag_id" xorm:"not null BIGINT(20) tag_id"`
	UserId     uint64    `json:"user_id" xorm:"not null BIGINT(20) user_id"`
	HappenedAt time.Time `json:"happened_at" xorm:"not null DATETIME happened_at"`
	CreatedAt  time.Time `json:"-" xorm:"created TIMESTAMP created_at"`
	UpdatedAt  time.Time `json:"-" xorm:"updated TIMESTAMP updated_at"`
}

func (i *Item) TableName() string {
	return "t_items"
}
