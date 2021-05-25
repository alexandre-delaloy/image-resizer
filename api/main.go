package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	setupServer()
}

func setupServer() *gin.Engine {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "[image-resizer]",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3010")
	return r
}