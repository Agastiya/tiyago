package utils

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB, name string, models any) {
	if !db.Migrator().HasTable(models) {
		if err := db.Migrator().CreateTable(models); err != nil {
			log.Fatal().Msgf("[Migration - %s] Failed: %v", name, err)
		}
		log.Info().Msgf("[Migration - %s] Successfully Migrate!", name)
	} else {
		log.Info().Msgf("[Migration - %s] Skipped (table already exists)", name)
	}
}

func DropTable(db *gorm.DB, name string, models any) {
	if db.Migrator().HasTable(models) {
		if err := db.Migrator().DropTable(models); err != nil {
			log.Fatal().Msgf("[Migration - %s] Failed: %v", name, err)
		}
		log.Info().Msgf("[Migration - %s] Successfully Migrate!", name)
	} else {
		log.Info().Msgf("[Migration - %s] Skipped (table not exists)", name)
	}
}

func AddColumn(db *gorm.DB, name string, models any, column string) {
	if !db.Migrator().HasColumn(models, column) {
		if err := db.Migrator().AddColumn(models, column); err != nil {
			log.Fatal().Msgf("[Migration - %s] Failed: %v", name, err)
		}
		log.Info().Msgf("[Migration - %s] Successfully migrated!", name)
	} else {
		log.Info().Msgf("[Migration - %s] Skipped (column already exists)", name)
	}
}

func DropColumn(db *gorm.DB, name string, models any, column string) {
	if db.Migrator().HasColumn(models, column) {
		if err := db.Migrator().DropColumn(models, column); err != nil {
			log.Fatal().Msgf("[Migration - %s] Failed: %v", name, err)
		}
		log.Info().Msgf("[Migration - %s] Successfully Migrate!", name)
	} else {
		log.Info().Msgf("[Migration - %s] Skipped (column not exists)", name)
	}
}
