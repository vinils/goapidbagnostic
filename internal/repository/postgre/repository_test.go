package postgre

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewRepositoryOpenConnection_WhenSuccess(test *testing.T) {
	var gormDb *gorm.DB = nil

	openSuccess := func(dialector gorm.Dialector, options ...gorm.Option) (*gorm.DB, error) {
		return gormDb, nil
	}

	expected := repo{
		database: gormDb,
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

func TestMigrateModels(test *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		test.Fatalf("failed to open sqlmock database: %s", err)
	}
	defer sqlDB.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db",
		DriverName:           "postgres",
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		test.Fatalf("failed to open gorm connection: %s", err)
	}

	// GORM first checks for the table's existence.
	// mock.ExpectQuery("SHOW TABLES LIKE ?").WillReturnRows(
	// 	sqlmock.NewRows([]string{"Tables_in_test_db", "categories"}).AddRow("categories"))

	repo := repo{
		database: gormDB,
	}

	err = repo.MigrateModels()

	assert.NotNil(test, err)
	err = mock.ExpectationsWereMet()
	assert.NoError(test, err, "Unmet expectations: %v", err)
}

func TestCategory(test *testing.T) {
	var gormDb *gorm.DB = nil

	repo := repo{
		database: gormDb,
	}

	actual := repo.Category()

	assert.Equal(test, NewCategory(gormDb), actual)
}
