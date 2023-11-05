package controller

import "github.com/kalougata/bookkeeping/internal/service"

type TagController struct {
	srv *service.TagService
}

func NewTagController(srv *service.TagService) *TagController {
	return &TagController{srv}
}
