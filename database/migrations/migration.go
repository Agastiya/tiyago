package migrations

import (
	"gorm.io/gorm"
)

func Up(db *gorm.DB) {
	UpCreateUsersTable(db)
}

func Down(db *gorm.DB) {}
