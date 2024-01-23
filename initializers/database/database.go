package database

import (
	"os"

	_ "github.com/awa-hc/backend/initializers/models"
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
	// DB.AutoMigrate(&models.TaskProject{}, &models.Reminders{}, &models.TaskProject{}, &models.TaskReminders{}, &models.Tasks{}, &models.UserProject{}, &models.Users{}, &models.Comment{}, &models.ProjectComment{})
}
