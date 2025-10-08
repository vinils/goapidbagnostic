package postgre

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewRepositoryOpenConnection_WhenSuccess(test *testing.T) {
	var gormDb *gorm.DB = nil

	openSuccess := func(dialector gorm.Dialector, options ...gorm.Option) (*gorm.DB, error) {
		return gormDb, nil
	}

	expected := repo{
		database: gormDb,
		Category: NewCategory(gormDb),
	}

	actual, err := newReposotiry("any", openSuccess)

	assert.Equal(test, &expected, actual)
	assert.Nil(test, err)
}

func TestNewRepositoryOpenConnection_WhenFail(test *testing.T) {
	errorMsg := "Generic Error"

	openFail := func(dialector gorm.Dialector, options ...gorm.Option) (*gorm.DB, error) {
		return nil, errors.New(errorMsg)
	}

	actual, err := newReposotiry("any", openFail)

	assert.Nil(test, actual)
	assert.NotNil(test, err)
	assert.EqualError(test, err, errorMsg)
}

func TestNewRepository_WhenFail(test *testing.T) {
	failCnnString := "any"
	actual, err := NewReposotiry(failCnnString)

	assert.Nil(test, actual)
	assert.NotNil(test, err)
}
