package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinils/goapitemplate/internal/entity"
	"github.com/vinils/goapitemplate/internal/repository"
)

type category struct{}

func sendError(ctx *gin.Context, code int, err error) {
	header := gin.H{"error": err.Error()}

	ctx.AbortWithStatusJSON(code, header)
}

func NewCategory() category { return category{} }

func (c category) Create(ctx *gin.Context, repo repository.ICategory) {
	var body struct {
		Name string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		return
	}

	category := entity.NewCategory(body.Name)

	if err := category.IsValid(); err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := repo.Create(category); err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, category)
}
