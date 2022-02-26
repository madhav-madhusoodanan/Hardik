package main

import (
	"madhav/hardik/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")

	server.GET("/:id", func(ctx *gin.Context) {

	})
	server.GET("/error", func(ctx *gin.Context) {
		ctx.HTML(400, "err.tmpl", nil)
	})
	server.GET("/delete/:key", handlers.Delete)
	server.GET("/modify/:key", handlers.Add)
	server.GET("/add/:key", handlers.Add)

	server.Run(":3000")
}