package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Avatar Avatar
}

type Users []User

type UserInput struct {
	Name string `json:"name"`
}
