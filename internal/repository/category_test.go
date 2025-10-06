package repository

import (
	"errors"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinils/goapitemplate/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type testData struct {
	database *gorm.DB
	mock     sqlmock.Sqlmock
}

var test testData

func TestMain(m *testing.M) {

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	// 2. Create a GORM DB instance using the mock connection and a dialector
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	test = testData{
		database: gormDB,
		mock:     mock,
	}

	// Run all tests
	code := m.Run()

	sqlDB.Close()
	os.Exit(code)
}

func TestNewCategory(test *testing.T) {
	var gormDb *gorm.DB = nil

	expected := category{db: gormDb}
	actual := NewCategory(gormDb)

	assert.Equal(test, actual, expected)
}

func TestCategoryCreate(t *testing.T) {
	// 3. Initialize your repository with the GORM DB instance
	repo := NewCategory(test.database)

	// 4. Define the expected SQL interactions
	category := entity.NewCategory("name")
	test.mock.ExpectBegin() // GORM often starts a transaction for create operations
	test.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "categories" ("name","created_at","updated_at") VALUES ($1,$2,$3)`)).
		WithArgs(category.Name, category.CreatedAt, category.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate a successful insert
	test.mock.ExpectCommit() // GORM commits the transaction

	// 5. Call the method under test
	newCategory, err := repo.Create(category)

	// 6. Assert the results and verify expectations
	assert.NoError(t, err)
	err = test.mock.ExpectationsWereMet()
	assert.NoError(t, err, "Unmet expectations: %v", err)
	assert.Equal(t, category, newCategory)
}

func TestCategoryList(t *testing.T) {
	// 3. Initialize your repository with the GORM DB instance
	repo := NewCategory(test.database)

	// 4. Define the expected SQL interactions
	expected := []entity.Category{
		{Name: "name1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "name2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	rows := sqlmock.NewRows([]string{"name", "created_at", "updated_at"}).
		AddRow(expected[0].Name, expected[0].CreatedAt, expected[0].UpdatedAt).
		AddRow(expected[1].Name, expected[1].CreatedAt, expected[1].UpdatedAt)
	test.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).
		WillReturnRows(rows)

	// 5. Call the method under test
	actual, err := repo.List()

	// 6. Assert the results and verify expectations
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
	err = test.mock.ExpectationsWereMet()
	assert.NoError(t, err, "Unmet expectations: %v", err)
}

func TestCategoryList_WhenError(t *testing.T) {
	// 3. Initialize your repository with the GORM DB instance
	repo := NewCategory(test.database)

	// 4. Define the expected SQL interactions
	expectedErr := errors.New("generic error")
	test.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories"`)).
		WillReturnError(expectedErr)

	// 5. Call the method under test
	actual, err := repo.List()

	// 6. Assert the results and verify expectations
	assert.Nil(t, actual)
	assert.EqualError(t, err, expectedErr.Error())
	err = test.mock.ExpectationsWereMet()
	assert.NoError(t, err, "Unmet expectations: %v", err)
}
