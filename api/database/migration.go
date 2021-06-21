package database

import (
	"github.com/blyndusk/image-resizer/api/models"
	"github.com/sirupsen/logrus"
)

func Migrate() {
	_ = Db.AutoMigrate(&models.User{}, &models.Avatar{})
	logrus.Info("Migrations done !")
}
