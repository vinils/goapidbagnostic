package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

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

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	// Run all tests
	code := m.Run()

	os.Exit(code)
}

func TestNewCategory(test *testing.T) {

	expected := category{}
	actual := NewCategory()

	assert.Equal(test, actual, expected)
}

func getMockContext() (*gin.Context, *httptest.ResponseRecorder) {
	res := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(res)

	return ctx, res
}

func getPostMockContext(body *bytes.Buffer) (*gin.Context, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(http.MethodPost, "/any", body)
	req.Header.Set("Content-Type", "application/json")

	ctx, res := getMockContext()
	ctx.Request = req

	return ctx, res
}

type requestBody struct {
	Name string
}

func (c requestBody) CastToByteBuffer() *bytes.Buffer {
	json, _ := json.Marshal(c)
	return bytes.NewBuffer(json)
}

func TestCreateCategory(t *testing.T) {
	expectedStatus := http.StatusCreated

	requestBody := requestBody{Name: "anyname"}
	ctx, response := getPostMockContext(requestBody.CastToByteBuffer())
	repo := repositoryMock{create: func(c entity.Category) error { return nil }}

	NewCategory().Create(ctx, repo)

	var responseCategory entity.Category
	err := json.Unmarshal(response.Body.Bytes(), &responseCategory)
	assert.NoError(t, err)
	assert.Equal(t, requestBody.Name, responseCategory.Name)
	assert.Equal(t, expectedStatus, response.Code)
}

func TestCreateCategory_WhenInvalidRequired(t *testing.T) {
	expectedErrorStatus := http.StatusBadRequest

	invalidCategoryName := ""
	body := requestBody{Name: invalidCategoryName}.CastToByteBuffer()
	ctx, response := getPostMockContext(body)
	repo := repositoryMock{create: func(c entity.Category) error { return nil }}

	NewCategory().Create(ctx, repo)

	assert.Equal(t, expectedErrorStatus, response.Code)
	assert.Contains(t, response.Body.String(), "name is required")
}

func TestCreateCategory_WhenCreateError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedErrorStatus := http.StatusBadRequest
	expectedErrorMsg := "any error"
	expectedError := errors.New(expectedErrorMsg)

	body := requestBody{Name: "anyname"}.CastToByteBuffer()
	ctx, response := getPostMockContext(body)
	repo := repositoryMock{create: func(c entity.Category) error { return expectedError }}

	NewCategory().Create(ctx, repo)

	assert.Equal(t, expectedErrorStatus, response.Code)
	assert.Contains(t, response.Body.String(), expectedErrorMsg)
}

func TestCreateCategory_WhenJSONBindError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedErrorStatus := http.StatusBadRequest

	errorJsonBindPayload := `{"name": "anyname"`
	ctx, response := getPostMockContext(bytes.NewBufferString(errorJsonBindPayload))

	NewCategory().Create(ctx, nil)

	assert.Equal(t, expectedErrorStatus, response.Code)
	assert.Contains(t, response.Body.String(), "unexpected EOF")
}
