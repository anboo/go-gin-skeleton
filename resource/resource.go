package resource

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Resource struct {
	Env Env
	DB  *gorm.DB
}

func Get(ctx context.Context) (*Resource, error) {
	r := Resource{}

	var err error

	err = r.initEnv()
	if err != nil {
		return nil, fmt.Errorf("init env: %w", err)
	}

	err = r.initDb()
	if err != nil {
		return nil, fmt.Errorf("init db: %w", err)
	}

	return &r, nil
}
