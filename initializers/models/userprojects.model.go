package models

import "gorm.io/gorm"

type UserProject struct {
	gorm.Model
	UserID    uint
	ProjectID uint
	User      Users    `gorm:"foreignKey:UserID"`
	Project   Projects `gorm:"foreignKey:ProjectID"`
}
