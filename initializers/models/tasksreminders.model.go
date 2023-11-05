package models

import "gorm.io/gorm"

type TaskReminders struct {
	gorm.Model
	ReminderID  uint
	ProjectID   uint
	TaskID      uint
	CreatedByID uint
	Reminder    Reminders `gorm:"foreignKey:ReminderID"`
	Project     Projects  `gorm:"foreignKey:ProjectID"`
	Task        Tasks     `gorm:"foreignKey:TaskID"`
	User        Users     `gorm:"foreignKey:CreatedByID"`
}
