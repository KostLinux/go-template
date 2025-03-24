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

func NewPostgreSQL(cfg *config.DatabaseConfig) Database {
	return &PostgreSQL{cfg: cfg}
}

func (p *PostgreSQL) Connect() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.cfg.Host, p.cfg.Port, p.cfg.User, p.cfg.Password, p.cfg.Name, p.cfg.SSLMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	p.db = db
	return nil
}

func (p *PostgreSQL) Close() error {
	return p.db.Close()
}

func (p *PostgreSQL) Ping(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

func (p *PostgreSQL) GetDB() *sqlx.DB {
	return p.db
}
