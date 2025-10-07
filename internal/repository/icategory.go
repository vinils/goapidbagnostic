package repository

import "github.com/vinils/goapitemplate/internal/entity"

type ICategory interface {
	Create(entity.Category) (entity.Category, error)
	List() ([]entity.Category, error)
}
