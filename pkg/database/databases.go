package database

import (
	"fmt"
	"go-template/config"
)

func New(cfg *config.DatabaseParams) (Database, error) {
	switch cfg.Driver {
	case "postgresql":
		return NewPostgreSQL(cfg), nil
	case "mysql":
		return NewMySQL(cfg), nil
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}
}
