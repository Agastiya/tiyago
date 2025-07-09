package migrations

import (
	"github.com/agastiya/tiyago/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func DropColumnStatusInUsersTable(db *gorm.DB) {
	if db.Migrator().HasColumn(&models.User{}, "status") {
		if err := db.Migrator().DropColumn(models.User{}, "status"); err != nil {
			log.Fatal().Msgf("[Migration - DropColumnStatusInUsersTable] Failed: %v", err)
		}
		log.Info().Msg("[Migration - DropColumnStatusInUsersTable] Successfully Migrate!")
	} else {
		log.Info().Msg("[Migration - DropColumnStatusInUsersTable] Skipped (column not exists)")
	}
}
