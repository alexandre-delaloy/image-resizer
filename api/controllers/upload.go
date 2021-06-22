package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	extension := filepath.Ext(file.Filename)
	if extension != ".png" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Only PNG files are supported",
		})
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	newFile := "./worker/images/uploaded/user_" + c.Params.ByName("id") + extension

	if err := c.SaveUploadedFile(file, newFile); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
	// worker.Produce()
	// worker.Consume()
}
