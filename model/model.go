package model

import (
	"template/entities"

	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

type Model interface {
	// Add new methods here
	GetUser(email string) (user *entities.User, err error)
}

func NewModel(db *gorm.DB) Model {
	return &model{
		db: db,
	}
}
