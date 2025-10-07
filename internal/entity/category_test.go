package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(test *testing.T) {
	myname := "myname"
	expected := Category{Name: myname}
	actual := NewCategory(myname)

	assert.Equal(test, actual.Name, expected.Name)
}

func TestNewCategoryIsValid_WhenNotValid(test *testing.T) {
	myname := ""
	actual := NewCategory(myname)
	err := actual.IsValid()

	assert.NotNil(test, err, "Expected an error, but got nil")
	assert.EqualError(test, err, "name is required", "Expected a specific error message")
}

func TestNewCategoryIsValid_WhenIsValid(test *testing.T) {
	myname := "lengthBiggerThan1"
	actual := NewCategory(myname)
	err := actual.IsValid()

	assert.Nil(test, err, "Expected nil, but got error")
}

func TestNewCategoryIsValid_WhenSmallerThan3(test *testing.T) {
	nameSmallerthan3 := "b"
	actual := NewCategory(nameSmallerthan3)
	err := actual.IsValid()

	assert.NotNil(test, err)
	assert.EqualError(test, err, "lenght name has to be bigger than 2")
}
