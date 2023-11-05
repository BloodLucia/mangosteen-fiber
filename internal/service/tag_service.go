package service

import "github.com/kalougata/bookkeeping/internal/data"

type TagService struct {
	data *data.Data
}

func NewTagService(data *data.Data) *TagService {
	return &TagService{data}
}
