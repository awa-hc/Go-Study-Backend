package models

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	Text        string
	ParentID    uint
	CreatedByID uint
	User        Users `gorm:"foreignKey:CreatedByID"`
}
