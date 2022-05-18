package database

import (
	"fmt"

	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *configs.Config) (DatabaseInterface, error) {
	connection, err := startDatabaseConnection(config)
	if err != nil {
		return nil, err
	}

	return &Database{connection: connection}, nil
}

func startDatabaseConnection(config *configs.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DatabaseConfig.Host,
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Name,
		config.DatabaseConfig.Port,
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}

type DatabaseInterface interface {
	GetConnection() *gorm.DB
}

type Database struct {
	connection *gorm.DB
}

func (d *Database) GetConnection() *gorm.DB {
	return d.connection
}
