package migrations

import (
	"github.com/agastiya/tiyago/config"
	"github.com/rs/zerolog/log"
)

func Up() error {
	db := config.DATABASE_MAIN.Get()
	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}
	log.Info().Msg("[Migration] Successfully Running!")
	return nil
}
