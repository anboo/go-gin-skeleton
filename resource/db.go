package resource

import (
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (r *Resource) initDb() error {
	log.Info().Msg("init db")
	defer log.Info().Msg("done init db")

	db, err := gorm.Open(postgres.Open(r.Env.DatabaseDSN), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("gorm open conn: %w", err)
	}

	r.DB = db

	err = goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	sqlDatabase, err := r.DB.DB()
	if err != nil {
		return fmt.Errorf("open sql db: %w", err)
	}

	err = goose.Up(sqlDatabase, "migration")
	if err != nil {
		return fmt.Errorf("do migration: %w", err)
	}

	return nil
}
