package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repository struct {
	connectionConfig ConnectionConfig
	database         *gorm.DB
}

// refactory review this code may not be here
type ConnectionConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func NewReposotiry(cnnConfig ConnectionConfig) (*repository, error) {
	return newReposotiry(cnnConfig, gorm.Open)
}

type openConnection func(gorm.Dialector, ...gorm.Option) (*gorm.DB, error)

func newReposotiry(cnnConfig ConnectionConfig, openConnection openConnection) (*repository, error) {
	cnnString := getConnectionString(cnnConfig)

	db, err := openConnection(postgres.Open(cnnString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	repo := &repository{
		database:         db,
		connectionConfig: cnnConfig,
	}

	return repo, err
}

func getConnectionString(config ConnectionConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		config.Host, config.Port, config.User, config.DBName, config.Password)
}
