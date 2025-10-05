package repository

import (
	"github.com/vinils/goapitemplate/internal/entity"
	"gorm.io/gorm"
)

type category struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) category {
	return category{db: db}
}

func (r category) Create(category entity.Category) error {
	return r.db.Create(&category).Error
}

func (r category) List() ([]entity.Category, error) {
	var categories []entity.Category

	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
