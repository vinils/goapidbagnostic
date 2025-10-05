package entity

import (
	"fmt"
	"time"
)

type Category struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) Category {
	category := Category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return category
}

func (c Category) IsValid() error {
	if len(c.Name) < 1 {
		return fmt.Errorf("name is required")
	}

	return nil
}
