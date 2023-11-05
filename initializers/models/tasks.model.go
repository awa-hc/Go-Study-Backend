package models

import (
	"time"

	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	Title       string
	Description string
	StartDate   time.Time
	TestDate    time.Time
	EndDate     time.Time
	Status      string
	Priority    int
	Tags        string
	CreatedByID uint
	User        Users `gorm:"foreignKey:CreatedByID"`
	ProjectID   uint
	Project     Projects `gorm:"foreingKey:ProjectID"`
}
