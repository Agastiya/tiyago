package migrations

import (
	"os"

	"github.com/agastiya/tiyago/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var modelList = []any{
	models.User{},
}

func Run(db *gorm.DB) {
	if err := db.AutoMigrate(modelList...); err != nil {
		log.Fatal().Msgf("[Migration] Failed: %v", err)
	}
	log.Info().Msg("[Migration] Successfully Running!")
	os.Exit(1)
}
