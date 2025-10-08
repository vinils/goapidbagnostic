package inmemory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepository(test *testing.T) {
	cat := NewCategory()
	expected := repo{
		category: &cat,
	}
	actual := NewRepository()

	assert.Equal(test, expected, actual)
}

func TestNewRepositoryCategory(test *testing.T) {
	repo := NewRepository()
	actual := repo.Category()

	assert.Equal(test, repo.category, actual)
}
