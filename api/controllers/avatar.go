package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/image-resizer/api/middlewares"
	"github.com/blyndusk/image-resizer/api/models"
)

func CreateAvatar(c *gin.Context) {
	var input models.AvatarInput
	middlewares.CreateAvatar(c, &input)
	c.JSON(http.StatusOK, input)
}

func GetAllAvatars(c *gin.Context) {
	var avatars models.Avatars
	middlewares.GetAllAvatars(c, &avatars)
	c.JSON(http.StatusOK, avatars)
}

func GetAvatarById(c *gin.Context) {
	var avatar models.Avatar
	middlewares.GetAvatarById(c, &avatar)
	c.JSON(http.StatusOK, avatar)
}

func UpdateAvatar(c *gin.Context) {
	var avatar models.Avatar
	var input models.AvatarInput
	middlewares.UpdateAvatar(c, &avatar, &input)
	c.JSON(http.StatusOK, avatar)
}

func DeleteAvatar(c *gin.Context) {
	var avatar models.Avatar
	middlewares.DeleteAvatar(c, &avatar)
	c.JSON(http.StatusOK, gin.H{
		"message": "Avatar deleted successfully",
	})
}
