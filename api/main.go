package main

import (
	"net/http"

	"github.com/blyndusk/image-resizer/api/database"
	"github.com/blyndusk/image-resizer/api/router"
	"github.com/blyndusk/image-resizer/api/queue/create_user"
	"github.com/gin-gonic/gin"
)

func main() {
	setupServer()
}

func setupServer() *gin.Engine {
	database.Connect()
	database.Migrate()

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

	r.GET("/emit", func(c * gin.Context) {
		create_user.emitCreateUser()
		c.JSON(http.StatusOK, gin.H {
			"message" : "le emit du create user",
		})
	})
	router.Setup(r)

	r.Run(":3010")
	return r
}

