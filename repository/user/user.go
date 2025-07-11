package user

import (
	"github.com/agastiya/tiyago/models"
)

type IUserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.PostgreDB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
