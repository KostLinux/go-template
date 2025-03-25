package database

import (
	"fmt"
	"go-template/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	ConnectionManager
}

func NewPostgreSQL(cfg *config.DatabaseParams) Database {
	return &PostgreSQL{
		ConnectionManager: ConnectionManager{cfg: cfg},
	}
}

func (pg *PostgreSQL) Connect() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		pg.cfg.Host, pg.cfg.User, pg.cfg.Password, pg.cfg.Name, pg.cfg.Port, pg.cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	pg.db = db
	return pg.SetConnectionParams()
}
