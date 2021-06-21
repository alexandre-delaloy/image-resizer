package middlewares

import (
	"net/http"

	"github.com/blyndusk/image-resizer/api/database"
	"github.com/blyndusk/image-resizer/api/helpers"
	"github.com/blyndusk/image-resizer/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateAvatar(c *gin.Context, input *models.AvatarInput) {
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	avatar := hydrateAvatar(input)
	if err := database.Db.Create(&avatar).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetAllAvatars(c *gin.Context, avatars *models.Avatars) {
	if err := database.Db.Preload("Avatar").Find(&avatars).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetAvatarById(c *gin.Context, avatar *models.Avatar) {
	if err := database.Db.Preload("Avatar").Where("id = ?", c.Params.ByName("id")).First(&avatar).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func UpdateAvatar(c *gin.Context, avatar *models.Avatar, input *models.AvatarInput) {
	GetAvatarById(c, avatar)
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	updatedAvatar := hydrateAvatar(input)
	database.Db.Model(&avatar).Updates(updatedAvatar)
}

func DeleteAvatar(c *gin.Context, avatar *models.Avatar) {
	if err := database.Db.First(&avatar, c.Params.ByName("id")).Delete(&avatar).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func hydrateAvatar(input *models.AvatarInput) models.Avatar {
	return models.Avatar{
		FilePath: input.FilePath,
	}
}
