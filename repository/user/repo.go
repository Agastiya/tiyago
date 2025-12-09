package user

import (
	"github.com/agastiya/tiyago/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	PostgreDB *gorm.DB
}

func New(db *gorm.DB) IUserRepository {
	return &UserRepository{PostgreDB: db}
}

type BrowseUserFilter struct {
	PageSize   int
	PageNumber int
	SortColumn string
	SortOrder  string
	Fullname   *string
	Username   *string
	Email      *string
}

type BrowseUserWithMeta struct {
	models.User
	TotalRecords int
	HasReachMax  bool
}
