package models

import (
	"time"

	"gorm.io/gorm"
)

type Reminders struct {
	gorm.Model
	ReminderDate time.Time
	Title        string
	Description  string
}
