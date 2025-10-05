package repository

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestConnectionString(test *testing.T) {
	host := "host"
	port := 46554
	user := "user"
	dbName := "dbname"
	password := "password"

	expected := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		host, port, user, dbName, password)

	cnnConfig := ConnectionConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DBName:   dbName,
		Password: password,
	}

	actual := getConnectionString(cnnConfig)

	assert.Equal(test, expected, actual)
}

func TestNewRepositoryOpenConnection_WhenSuccess(test *testing.T) {
	host := "host"
	port := 46554
	user := "user"
	dbName := "dbname"
	password := "password"

	cnnConfig := ConnectionConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DBName:   dbName,
		Password: password,
	}

	var gormDb *gorm.DB = nil

	openSuccess := func(dialector gorm.Dialector, options ...gorm.Option) (*gorm.DB, error) {
		return gormDb, nil
	}

	expected := repository{
		database:         gormDb,
		connectionConfig: cnnConfig,
	}

	actual, err := newReposotiry(cnnConfig, openSuccess)

	assert.Equal(test, &expected, actual)
	assert.Nil(test, err)
}

func TestNewRepositoryOpenConnection_WhenFail(test *testing.T) {
	host := "host"
	port := 46554
	user := "user"
	dbName := "dbname"
	password := "password"

	cnnConfig := ConnectionConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DBName:   dbName,
		Password: password,
	}

	errorMsg := "Generic Error"

	openFail := func(dialector gorm.Dialector, options ...gorm.Option) (*gorm.DB, error) {
		return nil, errors.New(errorMsg)
	}

	actual, err := newReposotiry(cnnConfig, openFail)

	assert.Nil(test, actual)
	assert.NotNil(test, err)
	assert.EqualError(test, err, errorMsg)
}

func TestNewRepository_WhenFail(test *testing.T) {
	host := "host"
	port := 46554
	user := "user"
	dbName := "dbname"
	password := "password"

	cnnConfig := ConnectionConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DBName:   dbName,
		Password: password,
	}

	actual, err := NewReposotiry(cnnConfig)

	assert.Nil(test, actual)
	assert.NotNil(test, err)
}
