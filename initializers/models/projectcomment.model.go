package models

import "gorm.io/gorm"

type ProjectComment struct {
	gorm.Model
	UserID    uint
	ProjectID uint
	CommentID uint
	Comment   Comment  `gorm:"foreignKey:CommentID"`
	User      Users    `gorm:"foreignKey:UserID"`
	Project   Projects `gorm:"foreignKey:ProjectID"`
}
