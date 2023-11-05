package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	cs := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(cs), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	// DB.AutoMigrate(&models.Comments{}, &models.TaskProject{}, &models.Reminders{}, &models.TaskProject{}, &models.TaskReminders{}, &models.Tasks{}, &models.UserProject{}, &models.Users{})
}
