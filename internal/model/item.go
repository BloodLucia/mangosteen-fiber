package model

import "time"

type Item struct {
	ID        uint64    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt time.Time `xorm:"deleted DATETIME deleted_at"`
	Amount    int
	TagId     int
	UserId    int
}

func (i *Item) TableName() string {
	return "t_items"
}

type ItemInReq struct {
	Amount int `json:"amount" validate:"int"`
	TagId  int `json:"tagId" validate:"int"`
	UserId int `json:"userId" validate:"int"`
}

func (u *ItemInReq) ToModel() *Item {
	return &Item{
		Amount: u.Amount,
		TagId:  u.TagId,
		UserId: u.UserId,
	}
}
