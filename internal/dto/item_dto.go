package dto

import (
	"github.com/gookit/goutil"
	"github.com/kalougata/bookkeeping/internal/model"
	"time"
)

type ItemInBody struct {
	Amount     int       `json:"amount" validate:"required|int" message:"int:{field} 必填且必须是数字类型"`
	Kind       string    `json:"kind" validate:"required|enum:income,expenses" message:"required:{field} 必填|enum:{field} 必须是income或者expenses"`
	TagId      string    `json:"tag_id" validate:"required" message:"required:{field} 必填"`
	UserId     string    `json:"user_id" validate:"required" message:"required:{field} 必填"`
	HappenedAt time.Time `json:"-" validate:"-"`
}

type ItemListQueries struct {
	HappenedAfter  string `query:"happened_after"`
	HappenedBefore string `query:"happened_before"`
	Page           int    `query:"page"`
	Limit          int    `query:"limit"`
	UserId         string `query:"-"`
}

func (r *ItemInBody) ToModel() *model.Item {
	return &model.Item{
		Amount:     r.Amount,
		Kind:       r.Kind,
		TagId:      goutil.Uint(r.TagId),
		UserId:     goutil.Uint(r.UserId),
		HappenedAt: r.HappenedAt,
	}
}
