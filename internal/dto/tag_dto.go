package dto

import (
	"github.com/gookit/goutil"
	"github.com/kalougata/bookkeeping/internal/model"
)

type TagInBody struct {
	Name   string `json:"name" validate:"required" message:"required:{field} 必填"`
	Kind   string `json:"type" validate:"required|enum:income,expenses" message:"required:{field} 必填|enum:{field} 必须是income或者expenses"`
	Sign   string `json:"sign" validate:"required" message:"required:{field} 必填"`
	UserId string `json:"userId" validate:"required" message:"required:{field} 必填"`
}

type TagListQueries struct {
	Kind   string `query:"kind"`
	Page   int    `query:"page"`
	UserId string `query:"-"`
}

func (body *TagInBody) ToModel() *model.Tag {
	return &model.Tag{
		Name:   body.Name,
		Kind:   body.Kind,
		Sign:   body.Sign,
		UserId: goutil.Uint(body.UserId),
	}
}
