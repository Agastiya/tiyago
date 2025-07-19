package migrations

import (
	"github.com/agastiya/tiyago/models"
	"gorm.io/gorm"
)

func CreateUsersTable(db *gorm.DB) {
	CreateTable(db, "CreateUsersTable", models.User{})
}
