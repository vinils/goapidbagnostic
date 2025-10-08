package inmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinils/goapitemplate/internal/entity"
)

func TestNewCategory(test *testing.T) {
	expected := category{
		db: &[]entity.Category{},
	}
	actual := NewCategory()

	assert.Equal(test, actual, expected)
}

func TestCategoryCreate(t *testing.T) {
	repo := NewCategory()
	category := entity.NewCategory("name")

	_, err := repo.Create(category)

	assert.NoError(t, err)
	assert.NoError(t, err, "Unmet expectations: %v", err)
	assert.Equal(t, category, (*repo.db)[0])
}

func TestCategoryList(t *testing.T) {
	repo := NewCategory()
	repo.db = &[]entity.Category{
		entity.NewCategory("name1"),
		entity.NewCategory("name2"),
	}

	// 5. Call the method under test
	actual, err := repo.List()

	// 6. Assert the results and verify expectations
	assert.NoError(t, err)
	assert.Equal(t, *repo.db, actual)
	assert.NoError(t, err, "Unmet expectations: %v", err)
}
