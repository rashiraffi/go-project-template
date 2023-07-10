package service

import (
	"template/entities"
	"template/model"
)

type service struct {
	model model.Model
}

type Service interface {
	// Add new methods here
	Login(userCred *entities.User) (token string, err error)
}

func NewService(m model.Model) Service {
	return &service{
		model: m,
	}
}
