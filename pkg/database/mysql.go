package database

import (
	"fmt"
	"go-template/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	ConnectionManager
}

func NewMySQL(cfg *config.DatabaseParams) Database {
	return &MySQL{
		ConnectionManager: ConnectionManager{cfg: cfg},
	}
}

func (conn *MySQL) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conn.cfg.User, conn.cfg.Password, conn.cfg.Host, conn.cfg.Port, conn.cfg.Name)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to mysql: %w", err)
	}

	conn.db = db
	return conn.SetConnectionParams()
}
