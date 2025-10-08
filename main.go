package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vinils/goapitemplate/internal/controller"
	"github.com/vinils/goapitemplate/internal/repository"
	"github.com/vinils/goapitemplate/internal/repository/inmemory"
	"github.com/vinils/goapitemplate/internal/repository/postgre"
)

type timeStruct struct {
	Time time.Time `json:"time"`
}

func newTimeNow() timeStruct { return newTime(time.Now()) }

func newTime(time time.Time) timeStruct { return timeStruct{time} }

func getHealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, newTimeNow())
}

func getRepository(cnnStr string) repository.IRepository {
	if cnnStr == "" {
		return inmemory.NewRepository()
	}

	ps, err := postgre.NewRepository(cnnStr)
	if err != nil {
		panic("Could not connect or invalid connection string" + err.Error())
	}

	if mgErr := ps.MigrateModels(); mgErr != nil {
		panic("Could not migrate model " + mgErr.Error())
	}

	return ps
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	cnnStr := os.Getenv("CNNSTR")
	repo := getRepository(cnnStr)
	ctrl := controller.NewCategory()
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.GET("/healthcheck", getHealthCheck)
		v1.POST("/category", func(ctx *gin.Context) {
			ctrl.Create(ctx, repo.Category())
		})
		v1.GET("/categories", func(ctx *gin.Context) {
			ctrl.List(ctx, repo.Category())
		})
	}
	return router
}

func main() {
	server := setupRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	_ = server.Run(":" + port)
}
