package models

import (
	"gorm.io/gorm"
)

type Avatar struct {
	gorm.Model
	FilePath string
	UserID   uint
}

type Avatars []Avatar

type AvatarInput struct {
	FilePath string `json:"file_path"`
}
