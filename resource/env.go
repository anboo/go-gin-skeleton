package resource

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Env struct {
	HttpPort    int    `env:"HTTP_PORT"`
	DatabaseDSN string `env:"DB_DSN"`
}

func (r *Resource) initEnv() error {
	log.Info().Msg("init env")
	defer log.Info().Msg("done init env")

	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		return fmt.Errorf("dot env read: %w", err)
	}

	if err := env.ParseWithOptions(&r.Env, env.Options{Environment: envs}); err != nil {
		return fmt.Errorf("env parse: %w", err)
	}
	return nil
}
