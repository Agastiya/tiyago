package user

import (
	"github.com/agastiya/tiyago/models"
	"gorm.io/gorm"
)

type (
	UserRepository struct {
		PostgreDB *gorm.DB
	}

	IUserRepository interface {
		CreateUser(user *models.User) (*models.User, error)
	}
)

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{PostgreDB: db}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.PostgreDB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
