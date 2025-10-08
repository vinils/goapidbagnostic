package inmemory

import (
	"github.com/vinils/goapitemplate/internal/entity"
	"github.com/vinils/goapitemplate/internal/repository"
)

type category struct {
	db *[]entity.Category
}

// Ensure that category implements the ICategory interface.
var _ repository.ICategory = category{}

func NewCategory() category {
	return category{
		db: &[]entity.Category{},
	}
}

func (r category) Create(category entity.Category) (entity.Category, error) {
	*r.db = append(*r.db, category)

	return category, nil
}

func (r category) List() ([]entity.Category, error) {
	return *r.db, nil
}
