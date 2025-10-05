package entity

import (
	"fmt"
	"time"
)

type category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) (*category, error) {
	category := &category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return category, nil
}

func (c category) IsValid() error {
	if len(c.Name) < 1 {
		return fmt.Errorf("name is required")
	}

	return nil
}
