package models

import (
	"gorm.io/gorm"
)

type Avatar struct {
	gorm.Model
	FilePath string `json:"file_path"`
	UserID   uint   `json:"user_id"`
}

type Avatars []Avatar

type AvatarInput struct {
	FilePath string `json:"file_path"`
}
