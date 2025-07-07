package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	DBConfigName string
)

const (
	DATABASE_MAIN DBConfigName = "postgreSql"
)

var dbConfigConnection = map[DBConfigName]*gorm.DB{}

func (d DBConfigName) Get() *gorm.DB {
	return dbConfigConnection[d]
}

func CreatePostgreSQLConnection(env DatabaseSetting) *gorm.DB {
	psqlUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", env.Host, env.Port, env.User, env.Password, env.Database)
	connection, err := gorm.Open(postgres.Open(psqlUrl), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal().Msgf("[PostgreSQL] : failed to connect database. Err : %v", err)
	}

	log.Info().Msg("[PostgreSQL] Successfully Connected!")
	return connection
}
