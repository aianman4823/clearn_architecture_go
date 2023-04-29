package middlewares

import (
	"github.com/aianman4823/clearn_architecture_go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello, World")
	})
	router.GET("/tests", controllers.ListTest)
	router.GET("/tests/:id", controllers.GetTest)
}
