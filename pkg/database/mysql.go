package database

import (
	"context"
	"fmt"
	"go-template/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	db  *sqlx.DB
	cfg *config.DatabaseConfig
}

func NewMySQL(cfg *config.DatabaseConfig) Database {
	return &MySQL{cfg: cfg}
}

func (m *MySQL) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		m.cfg.User, m.cfg.Password, m.cfg.Host, m.cfg.Port, m.cfg.Name)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to mysql: %w", err)
	}

	m.db = db
	return nil
}

func (m *MySQL) Close() error {
	return m.db.Close()
}

func (m *MySQL) Ping(ctx context.Context) error {
	return m.db.PingContext(ctx)
}

func (m *MySQL) GetDB() *sqlx.DB {
	return m.db
}
