package main

import (
	"fmt"

	"github.com/aianman4823/clearn_architecture_go/internal/config"
	"github.com/aianman4823/clearn_architecture_go/internal/external"
	"github.com/aianman4823/clearn_architecture_go/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize databases
	external.SetupDB()

	// Setup webserver
	app := gin.Default()
	app.Use(middlewares.Transaction())
	app.Use(middlewares.Cors())
	middlewares.Routers(app)
	app.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
}
