package repository

import (
	"github.com/vinils/goapitemplate/internal/entity"
	"gorm.io/gorm"
)

type category struct {
	db *gorm.DB
}

// Ensure that category implements the ICategory interface.
var _ ICategory = category{}

func NewCategory(db *gorm.DB) category {
	return category{db: db}
}

func (r category) Create(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r category) List() ([]entity.Category, error) {
	var categories []entity.Category

	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
