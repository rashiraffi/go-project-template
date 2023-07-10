package model

import (
	"log"
	"template/entities"

	"gorm.io/gorm"
)

func (m *model) GetUser(email string) (user *entities.User, err error) {
	if err := m.db.Where("email=?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return
}
