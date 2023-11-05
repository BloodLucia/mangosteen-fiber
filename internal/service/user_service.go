package service

import "github.com/kalougata/bookkeeping/internal/data"

type UserService struct {
	data *data.Data
}

func (us *UserService) LoginWithEmail() error {
	return nil
}

func NewUserService(data *data.Data) *UserService {
	return &UserService{data}
}
