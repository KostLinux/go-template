package database

import (
	"context"
	"fmt"
	"go-template/config"

	"github.com/jmoiron/sqlx"
)

// Database represents a generic database interface
type Database interface {
	Connect() error
	Close() error
	Ping(ctx context.Context) error
	GetDB() *sqlx.DB
}

// NewConnection creates a new database connection based on driver
func NewConnection(cfg *config.DatabaseConfig) (Database, error) {
	switch cfg.Driver {
	case "postgres":
		return NewPostgreSQL(cfg), nil
	case "mysql":
		return NewMySQL(cfg), nil
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}
}
