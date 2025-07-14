package user

import (
	"github.com/agastiya/tiyago/models"
)

type IUserRepository interface {
	CheckUsernameExists(username string, id int64) (bool, error)
	CheckEmailExists(email string, id int64) (bool, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

func (r *UserRepository) CheckUsernameExists(username string, id int64) (bool, error) {
	var exists bool
	err := r.PostgreDB.
		Raw("SELECT EXISTS(SELECT 1 FROM users WHERE username = ? and id != ?)", username, id).
		Scan(&exists).Error
	return exists, err
}

func (r *UserRepository) CheckEmailExists(email string, id int64) (bool, error) {
	var exists bool
	err := r.PostgreDB.
		Raw("SELECT EXISTS(SELECT 1 FROM users WHERE email = ? and id != ?)", email, id).
		Scan(&exists).Error
	return exists, err
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.PostgreDB.Create(&user).Error
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.PostgreDB.Where("id = ?", user.Id).Updates(user).Error
}

func (r *UserRepository) DeleteUser(user *models.User) error {
	return r.PostgreDB.Where("id = ?", user.Id).Updates(&user).Error
}
