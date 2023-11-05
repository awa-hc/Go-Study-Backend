package models

import (
	"time"

	"gorm.io/gorm"
)

type Projects struct {
	gorm.Model
	Title       string
	Company     string
	Description string
	StartDate   time.Time `gorm:"required"`
	TestDate    time.Time
	EndDate     time.Time `gorm:"required"`
	CreatedByID uint
	Code        string `gorm:"unique"`
	Tags        string
}
