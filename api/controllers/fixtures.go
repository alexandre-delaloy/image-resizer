package controllers

import (
	"net/http"

	"github.com/blyndusk/image-resizer/api/database"
	"github.com/blyndusk/image-resizer/api/models"
	"github.com/gin-gonic/gin"
)

func LoadData(c *gin.Context) {
	users := models.Users{
		{
			Name:      "John Doe",
		},
	}

	database.Db.Create(&users)
	c.JSON(http.StatusOK, gin.H{
		"message": "Fixtures loaded",
	})

}
