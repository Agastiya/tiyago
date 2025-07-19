package migrations

import (
	"github.com/agastiya/tiyago/models"
	"gorm.io/gorm"
)

func DropColumnStatusInUsersTable(db *gorm.DB) {
	DropColumn(db, "DropColumnStatusInUsersTable", &models.User{}, "status")
}
