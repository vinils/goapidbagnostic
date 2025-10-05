package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinils/goapitemplate/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewCategory(test *testing.T) {
	var gormDb *gorm.DB = nil

	expected := category{db: gormDb}
	actual := NewCategory(gormDb)

	assert.Equal(test, actual, expected)
}

func TestCategoryCreate(t *testing.T) {
	// 1. Create a mock database connection using go-sqlmock
	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer sqlDB.Close()

	// 2. Create a GORM DB instance using the mock connection and a dialector
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	assert.NoError(t, err)

	// 3. Initialize your repository with the GORM DB instance
	repo := NewCategory(gormDB)

	// 4. Define the expected SQL interactions
	category := entity.NewCategory("name")
	mock.ExpectBegin() // GORM often starts a transaction for create operations
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "categories" ("name","created_at","updated_at") VALUES ($1,$2,$3)`)).
		WithArgs(category.Name, category.CreatedAt, category.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate a successful insert
	mock.ExpectCommit() // GORM commits the transaction

	// 5. Call the method under test
	err = repo.Create(&category)

	// 6. Assert the results and verify expectations
	assert.NoError(t, err)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Unmet expectations: %v", err)
}
