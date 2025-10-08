package inmemory

import "github.com/vinils/goapitemplate/internal/repository"

type repo struct {
	category *category
}

// Ensure that category implements the ICategory interface.
var _ repository.IRepository = repo{}

func NewRepository() repo {
	cat := NewCategory()
	return repo{
		category: &cat,
	}
}

func (r repo) Category() repository.ICategory {
	return r.category
}
