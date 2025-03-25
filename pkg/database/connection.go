package database

import (
	"context"
	"fmt"
	"time"

	"go-template/config"

	"gorm.io/gorm"
)

type Database interface {
	Connect() error
	Close() error
	Ping(ctx context.Context) error
	GetDB() *gorm.DB
}

type ConnectionManager struct {
	db  *gorm.DB
	cfg *config.DatabaseParams
}

func (cm *ConnectionManager) Close() error {
	sqlDB, err := cm.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	return sqlDB.Close()
}

func (cm *ConnectionManager) Ping(ctx context.Context) error {
	sqlDB, err := cm.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	return sqlDB.PingContext(ctx)
}

func (cm *ConnectionManager) GetDB() *gorm.DB {
	return cm.db
}

func (cm *ConnectionManager) SetConnectionParams() error {
	sqlDB, err := cm.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(cm.cfg.Connection.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cm.cfg.Connection.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cm.cfg.Connection.ConnMaxLifetime) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(cm.cfg.Connection.ConnMaxIdleTime) * time.Second)
	return nil
}
