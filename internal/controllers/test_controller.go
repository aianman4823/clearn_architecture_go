package controllers

import (
	"errors"

	"github.com/aianman4823/clearn_architecture_go/internal/repositories"
	"github.com/aianman4823/clearn_architecture_go/internal/usecases"
	"github.com/gin-gonic/gin"
)

func ListTest(ctx *gin.Context) {
	repository := repositories.NewTestRepository(DB(ctx))
	usecase := usecases.NewTestUsecase(repository)
	result, err := usecase.ListTest()
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("Not Found"))
	}
}

func GetTest(ctx *gin.Context) {
	repository := repositories.NewTestRepository(DB(ctx))
	usecase := usecases.NewTestUsecase(repository)
	result, err := usecase.GetTest(ctx.Param("id"))
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("Not Found"))
	}
}
