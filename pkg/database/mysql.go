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

func NewMySQL(cfg *config.DatabaseConfig) *MySQL {
	return &MySQL{cfg: cfg}
}

func (mysql *MySQL) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		mysql.cfg.User, mysql.cfg.Password, mysql.cfg.Host, mysql.cfg.Port, mysql.cfg.Name)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to mysql: %w", err)
	}

	mysql.db = db
	return nil
}

func (mysql *MySQL) Close() error {
	return mysql.db.Close()
}

func (mysql *MySQL) Ping(ctx context.Context) error {
	return mysql.db.PingContext(ctx)
}

func (mysql *MySQL) GetDB() *sqlx.DB {
	return mysql.db
}
