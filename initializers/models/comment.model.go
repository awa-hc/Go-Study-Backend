package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text          string
	ParentID      *uint
	CreatedByID   uint
	User          Users `gorm:"foreignKey:CreatedByID"`
	ProjectID     uint
	Project       Projects  `gorm:"foreignKey:ProjectID"`
	ChildComments []Comment `gorm:"foreignKey:ParentID"`
}
