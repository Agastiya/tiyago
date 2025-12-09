package auth

import "gorm.io/gorm"

type AuthRepository struct {
	PostgreDB *gorm.DB
}

func New(db *gorm.DB) IAuthRepository {
	return &AuthRepository{PostgreDB: db}
}
