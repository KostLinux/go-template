package database

import (
	"context"
	"fmt"

	"go-template/config"
	"go-template/model"

	"gorm.io/gorm"
)

type DBManager struct {
	db  *gorm.DB
	cfg *config.DatabaseConfig
}

func NewDBManager(cfg *config.DatabaseConfig) (*DBManager, error) {
	db, err := NewConnection(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create database connection: %w", err)
	}

	return &DBManager{
		db:  db,
		cfg: cfg,
	}, nil
}

func (m *DBManager) Connect() error {
	if err := m.db.AutoMigrate(&model.User{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}

func (m *DBManager) Close() error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (m *DBManager) Ping(ctx context.Context) error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}

func (m *DBManager) GetDB() *gorm.DB {
	return m.db
}
