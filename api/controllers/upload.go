package controllers

import (
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")
	extension := filepath.Ext(file.Filename)
	// if extension != "png" {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"message": "Only PNG files are supported",
	// 	})
	// 	return
	// }
	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)
	// Retrieve file information
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := strconv.Itoa(r1.Intn(100000)) + extension

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "./service-worker/producer/test-images/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}
