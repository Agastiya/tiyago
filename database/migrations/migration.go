package migrations

import "gorm.io/gorm"

func Run(db *gorm.DB) {
	CreateUsersTable(db)
	DropColumnStatusInUsersTable(db)
	AddColumnDeletedInUsersTable(db)
}
