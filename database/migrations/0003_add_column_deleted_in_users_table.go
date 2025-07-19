package migrations

import (
	"github.com/agastiya/tiyago/models"
	"gorm.io/gorm"
)

func AddColumnDeletedInUsersTable(db *gorm.DB) {
	AddColumn(db, "AddColumnDeletedInUsersTable", &models.User{}, "deleted_by")
	AddColumn(db, "AddColumnDeletedInUsersTable", &models.User{}, "deleted_at")
}
