package database

import (
	"fmt"

	"go-template/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	switch cfg.Driver {
	case "postgresql":
		return connectPostgres(cfg)
	case "mysql":
		return connectMySQL(cfg)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}
}

func connectPostgres(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func connectMySQL(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
