package postgre

import (
	"github.com/vinils/goapitemplate/internal/entity"
	"github.com/vinils/goapitemplate/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repo struct {
	database *gorm.DB
}

// Ensure that category implements the ICategory interface.
var _ repository.IRepository = repo{}

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
	}

	return repo, err
}

func (r *repo) MigrateModels() error {
	return r.database.AutoMigrate(&entity.Category{})
}

func (r repo) Category() repository.ICategory {
	return NewCategory(r.database)
}
