package migrations

import "github.com/agastiya/tiyago/config"

func Up() error {
	db := config.DATABASE_MAIN.Get()
	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}
	return nil
}
