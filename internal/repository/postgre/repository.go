package postgre

import (
	"github.com/vinils/goapitemplate/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repo struct {
	database *gorm.DB
	category repository.ICategory
}

// Ensure that category implements the ICategory interface.
var _ repository.ICategory = category{}

func NewReposotiry(cnnString string) (*repo, error) {
	return newReposotiry(cnnString, gorm.Open)
}

type openConnection func(gorm.Dialector, ...gorm.Option) (*gorm.DB, error)

func newReposotiry(cnnString string, openConnection openConnection) (*repo, error) {
	db, err := openConnection(postgres.Open(cnnString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	repo := &repo{
		database: db,
		category: NewCategory(db),
	}

	return repo, err
}
