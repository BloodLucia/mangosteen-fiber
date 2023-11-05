package model

import "time"

type Tag struct {
	ID        uint64    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt time.Time `xorm:"deleted DATETIME deleted_at"`
	Name      string
	Type      string
	Sign      string
	UserId    uint64
}

func (u *Tag) TableName() string {
	return "t_tags"
}

type TagInReq struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Sign   string `json:"sign"`
	UserId uint64 `json:"-"`
}

func (u *TagInReq) ToModel() *Tag {
	return &Tag{
		Name: u.Name,
		Type: u.Type,
		Sign: u.Sign,
	}
}
