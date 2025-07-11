package repository

import (
	"github.com/agastiya/tiyago/repository/user"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo user.IUserRepository
}

func InitRepos(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo: user.New(db),
	}
}
