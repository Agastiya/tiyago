package user

import "gorm.io/gorm"

type UserRepository struct {
	PostgreDB *gorm.DB
}

func New(db *gorm.DB) IUserRepository {
	return &UserRepository{PostgreDB: db}
}
