package handlers

import (
	"net/http"

	"github.com/awa-hc/backend/api/utils"
	"github.com/awa-hc/backend/initializers/database"
	"github.com/awa-hc/backend/initializers/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	user, _ := c.Get("user")
	userid := user.(models.Users).ID

	var taskprojects []models.TaskProject
	var task models.Tasks

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	task.CreatedByID = userid

	if userid == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user creator"})
		return
	}

	if task.ProjectID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project"})
		return
	}

	if !utils.ValidateRequiredStringFields(c, "title", task.Title) || !utils.ValidateRequiredStringFields(c, "status", task.Status) || !utils.ValidateRequiredStringFields(c, "tags", task.Tags) || !utils.ValidateRequiredStringFields(c, "description", task.Description) || !utils.ValidateRequiredTimeFields(c, "startDate", task.StartDate) || !utils.ValidateRequiredTimeFields(c, "endDate", task.EndDate) {
		return
	}

	if !utils.ValidateTimeFields(c, task.StartDate, task.TestDate, task.EndDate) {
		return
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creating task"})
		return
	}

	taskprojects = append(taskprojects, models.TaskProject{TaskID: task.ID, ProjectID: task.ProjectID})
	if err := database.DB.Create(&taskprojects).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creating task in taskprojects"})
		return
	}

	if err := database.DB.Preload("User").Preload("Project").Where("id = ?", task.ID).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error preloading taskprojects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": task})

}
func GetTasks(c *gin.Context) {
	var tasks []models.Tasks
	user, _ := c.Get("user")
	userid := user.(models.Users).ID

	if userid == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user creator"})
		return
	}

	database.DB.Preload("User").Preload("Project").Find(&tasks, "created_by_id = ?", userid)

	c.JSON(http.StatusOK, gin.H{"data": tasks})

}

func GetTask(c *gin.Context) {
	var task models.Tasks
	var id = c.Param("id")
	user, _ := c.Get("user")
	userid := user.(models.Users).ID

	if userid == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user creator"})
		return
	}
	if err := database.DB.Preload("User").Preload("Project").First(&task, "id = ? AND created_by_id = ?", id, userid).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})

}
