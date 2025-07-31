package migrations

import (
	"github.com/agastiya/tiyago/models"
	"github.com/agastiya/tiyago/pkg/helper/utils"
	"gorm.io/gorm"
)

func UpCreateUsersTable(db *gorm.DB) {
	utils.CreateTable(db, "Create table users", models.User{})
}

func DownCreateUsersTable(db *gorm.DB) {
	utils.DropTable(db, "Drop Table users", models.User{})
}
