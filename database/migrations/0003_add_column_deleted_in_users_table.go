package migrations

import (
	"github.com/agastiya/tiyago/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func AddColumnDeletedInUsersTable(db *gorm.DB) {
	// Add "deleted_by" column if it does not exist
	if !db.Migrator().HasColumn(&models.User{}, "deleted_by") {
		if err := db.Migrator().AddColumn(&models.User{}, "deleted_by"); err != nil {
			log.Fatal().Msgf("[Migration - AddColumnDeletedByInUsersTable] Failed: %v", err)
		}
		log.Info().Msg("[Migration - AddColumnDeletedByInUsersTable] Successfully migrated!")
	} else {
		log.Info().Msg("[Migration - AddColumnDeletedByInUsersTable] Skipped (column already exists)")
	}

	// Add "deleted_at" column if it does not exist
	if !db.Migrator().HasColumn(&models.User{}, "deleted_at") {
		if err := db.Migrator().AddColumn(&models.User{}, "deleted_at"); err != nil {
			log.Fatal().Msgf("[Migration - AddColumnDeletedAtInUsersTable] Failed: %v", err)
		}
		log.Info().Msg("[Migration - AddColumnDeletedAtInUsersTable] Successfully migrated!")
	} else {
		log.Info().Msg("[Migration - AddColumnDeletedAtInUsersTable] Skipped (column already exists)")
	}
}
