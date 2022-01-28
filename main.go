package main

import (
	"github.com/gin-gonic/gin"
	"github.com/techagent/ceeyuapp/handler"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./ui/html/*.gohtml")
	//router.StaticFS("static", http.Dir("./ui/static/"))

	router.GET("/", handler.Index())

	router.GET("/:id", handler.View())
	router.GET("/admin", handler.Admin())
	router.POST("/comment/count", handler.Counter())

	router.Run(":8080")
}