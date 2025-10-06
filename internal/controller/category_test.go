package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinils/goapitemplate/internal/entity"
)

type repositoryMock struct {
	create func(entity.Category) error
	list   func() ([]entity.Category, error)
}

func (r repositoryMock) List() ([]entity.Category, error) { return r.list() }
func (r repositoryMock) Create(c entity.Category) error   { return r.create(c) }

func TestNewCategory(test *testing.T) {

	expected := category{}
	actual := NewCategory()

	assert.Equal(test, actual, expected)
}

func TestCreateCategory(t *testing.T) {
	// 1. Set Gin to TestMode
	gin.SetMode(gin.TestMode)

	// 2. Create a new CategoryController instance
	controller := NewCategory()

	categoryName := "anyname"
	expectedStatus := http.StatusCreated
	expectedBody := entity.Category{Name: categoryName, CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	// Convert request body to JSON
	requestBody := entity.NewCategory(categoryName)
	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPost, "/any", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Create a test Gin context
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the controller method
	repo := repositoryMock{create: func(c entity.Category) error { return nil }}
	controller.Create(c, repo)

	// Assert the HTTP status code
	actualStatus := w.Code
	assert.Equal(t, expectedStatus, actualStatus)

	// Assert the response body
	var actualCategory entity.Category
	err := json.Unmarshal(w.Body.Bytes(), &actualCategory)
	assert.NoError(t, err)
	assert.Equal(t, expectedBody.Name, actualCategory.Name)
}

func TestCreateCategory_WhenInvalidRequired(t *testing.T) {
	// 1. Set Gin to TestMode
	gin.SetMode(gin.TestMode)

	// 2. Create a new CategoryController instance
	controller := NewCategory()

	invalidCategoryName := ""
	expectedErrorStatus := http.StatusBadRequest

	// Convert request body to JSON
	requestBody := entity.NewCategory(invalidCategoryName)
	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPost, "/any", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Create a test Gin context
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the controller method
	repo := repositoryMock{create: func(c entity.Category) error { return nil }}
	controller.Create(c, repo)

	// Assert the HTTP status code
	actualStatus := w.Code
	assert.Equal(t, expectedErrorStatus, actualStatus)

	// Assert the response body
	assert.Contains(t, w.Body.String(), "name is required")
}

func TestCreateCategory_WhenCreateError(t *testing.T) {
	// 1. Set Gin to TestMode
	gin.SetMode(gin.TestMode)

	// 2. Create a new CategoryController instance
	controller := NewCategory()

	expectedErrorStatus := http.StatusBadRequest
	expectedErrorMsg := "any error"
	expectedError := errors.New(expectedErrorMsg)

	// Convert request body to JSON
	requestBody := entity.NewCategory("anyname")
	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(http.MethodPost, "/any", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Create a test Gin context
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the controller method
	repo := repositoryMock{create: func(c entity.Category) error { return expectedError }}
	controller.Create(c, repo)

	// Assert the HTTP status code
	actualStatus := w.Code
	assert.Equal(t, expectedErrorStatus, actualStatus)

	// Assert the response body
	assert.Contains(t, w.Body.String(), expectedErrorMsg)
}

func TestCreateCategory_WhenJSONBindError(t *testing.T) {
	// 1. Set Gin to TestMode
	gin.SetMode(gin.TestMode)

	// 2. Create a new CategoryController instance
	controller := NewCategory()

	expectedErrorStatus := http.StatusBadRequest

	// Convert request body to JSON
	errorJsonBindPayload := `{"name": "anyname"`
	req, _ := http.NewRequest(http.MethodPost, "/any", bytes.NewBufferString(errorJsonBindPayload))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Create a test Gin context
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the controller method
	controller.Create(c, nil)

	// Assert the HTTP status code
	actualStatus := w.Code
	assert.Equal(t, expectedErrorStatus, actualStatus)

	// Assert the response body
	assert.Contains(t, w.Body.String(), "unexpected EOF")
}
