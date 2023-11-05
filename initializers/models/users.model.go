package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Fullname   string
	Email      string `gorm:"unique"`
	Password   string
	ImgProfile string
	Username   string `gorm:"unique"`
	UserRole   string
}
