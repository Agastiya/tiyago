package migrations

import (
	"github.com/agastiya/tiyago/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func CreateUsersTable(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.User{}) {
		if err := db.Migrator().CreateTable(models.User{}); err != nil {
			log.Fatal().Msgf("[Migration - CreateUsersTable] Failed: %v", err)
		}
		log.Info().Msg("[Migration - CreateUsersTable] Successfully Migrate!")
	} else {
		log.Info().Msg("[Migration - CreateUsersTable] Skipped (table already exists)")
	}
}
