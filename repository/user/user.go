package user

import (
	"github.com/agastiya/tiyago/models"
)

type IUserRepository interface {
	CheckUsernameExists(username string) (bool, error)
	CheckEmailExists(email string) (bool, error)
	CreateUser(user *models.User) (*models.User, error)
}

func (r *UserRepository) CheckUsernameExists(username string) (bool, error) {
	var exists bool
	err := r.PostgreDB.
		Raw("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).
		Scan(&exists).Error
	return exists, err
}

func (r *UserRepository) CheckEmailExists(email string) (bool, error) {
	var exists bool
	err := r.PostgreDB.
		Raw("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).
		Scan(&exists).Error
	return exists, err
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.PostgreDB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
