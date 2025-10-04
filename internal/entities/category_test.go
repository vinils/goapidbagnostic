package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(test *testing.T) {
	myname := "myname"
	expected := category{Name: myname}
	actual, err := NewCategory(myname)

	if err != nil {
		assert.Fail(test, "error while creating category")
	}

	assert.Equal(test, actual.Name, expected.Name)
}
