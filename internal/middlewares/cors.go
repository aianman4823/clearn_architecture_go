package middlewares

import (
	myConfig "github.com/aianman4823/clearn_architecture_go/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = myConfig.CorsAllowOrigin
	config.AllowCredentials = true
	return cors.New(config)
}
