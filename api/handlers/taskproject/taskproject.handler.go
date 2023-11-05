package handlers

import (
	"net/http"

	"github.com/awa-hc/backend/initializers/database"
	"github.com/awa-hc/backend/initializers/models"
	"github.com/gin-gonic/gin"
)

func GetTaskProject(c *gin.Context) {
	var id = c.Param("id")
	var taskproject []models.TaskProject

	if err := database.DB.
		Preload("Task").
		Preload("Task.User").
		Preload("Task.Project").
		Preload("Project").
		Where("project_id = ?", id).
		Find(&taskproject).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find taskproject"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"taskproject": taskproject})
}
