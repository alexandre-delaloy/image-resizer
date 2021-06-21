package controllers

import (
	"net/http"

	"github.com/blyndusk/image-resizer/api/database"
	"github.com/blyndusk/image-resizer/api/models"
	"github.com/gin-gonic/gin"
)

func LoadData(c *gin.Context) {

	user := models.User{
		Name:   "John Doe",
		Avatar: models.Avatar{FilePath: "/path/to/avatar"},
	}
	database.Db.Create(&user)
	c.JSON(http.StatusOK, user)

}
