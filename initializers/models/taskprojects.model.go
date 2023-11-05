package models

import "gorm.io/gorm"

type TaskProject struct {
	gorm.Model
	ProjectID uint
	TaskID    uint
	Project   Projects `gorm:"foreignKey:ProjectID"`
	Task      Tasks    `gorm:"foreignKey:TaskID"`
}
