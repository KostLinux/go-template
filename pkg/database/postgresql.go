package database

import (
	"context"
	"fmt"

	"go-template/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	db  *sqlx.DB
	cfg *config.DatabaseConfig
}

func NewPostgreSQL(cfg *config.DatabaseConfig) *PostgreSQL {
	return &PostgreSQL{cfg: cfg}
}

func (postgres *PostgreSQL) Connect() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		postgres.cfg.Host, postgres.cfg.Port, postgres.cfg.User, postgres.cfg.Password, postgres.cfg.Name, postgres.cfg.SSLMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	postgres.db = db
	return nil
}

func (postgres *PostgreSQL) Close() error {
	return postgres.db.Close()
}

func (postgres *PostgreSQL) Ping(ctx context.Context) error {
	return postgres.db.PingContext(ctx)
}

func (postgres *PostgreSQL) GetDB() *sqlx.DB {
	return postgres.db
}
