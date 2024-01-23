package handlers

import (
	"net/http"

	"github.com/awa-hc/backend/api/utils"
	"github.com/awa-hc/backend/initializers/database"
	"github.com/awa-hc/backend/initializers/models"
	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	var project models.Projects
	user, _ := c.Get("user")
	userid := user.(models.Users).ID

	project.CreatedByID = userid

	if userid == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user creator"})
		return
	}

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	if !utils.ValidateRequiredStringFields(c, "title", project.Title) || !utils.ValidateRequiredStringFields(c, "tags", project.Tags) || !utils.ValidateRequiredStringFields(c, "company", project.Company) || !utils.ValidateRequiredStringFields(c, "description", project.Description) || !utils.ValidateRequiredTimeFields(c, "startDate", project.StartDate) || !utils.ValidateRequiredTimeFields(c, "endDate", project.EndDate) {
		return
	}

	if !utils.ValidateTimeFields(c, project.StartDate, project.TestDate, project.EndDate) {
		return
	}

	if err := database.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create the project"})
		return
	}

	var userProject models.UserProject
	userProject.UserID = userid
	userProject.ProjectID = project.ID

	if err := database.DB.Create(&userProject).Error; err != nil {
		database.DB.Delete(&project)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create the project on userproject"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "project created successfully", "project": project})

}

func GetProjects(c *gin.Context) {
	var project []models.Projects
	if err := database.DB.Find(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find projects"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"projects": project})
}

func GetProject(c *gin.Context) {
	c.Param("id")
	var project models.Projects

	if err := database.DB.Where("id = ?", c.Param("id")).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"project": project})

}

func DeleteProject(c *gin.Context) {
	var id = c.Param("id")
	user, _ := c.Get("user")
	userid := user.(models.Users).ID

	var project models.Projects
	projecttitle := project.Title
	projectcreador := project.CreatedByID

	var input struct {
		Title        string `json:"title"`
		Confirmation string `json:"confirmation"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	if !utils.ValidateRequiredStringFields(c, "confirmation", input.Confirmation) {
		return
	}

	if err := database.DB.Where("id = ?", id).First(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find project"})
		return
	}

	if userid != projectcreador {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are not the creator of this project"})
		return
	}

	if projecttitle != input.Title {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong confirmation"})
		return
	}
	if input.Title != input.Confirmation {
		c.JSON(http.StatusBadRequest, gin.H{"error": "confirmation does not match"})
		return
	}

	if err := database.DB.Where("id = ?", id).Delete(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not delete project"})
		return
	}

	if err := database.DB.Where("project_id = ?", id).Delete(&models.UserProject{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not delete project on userproject table"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "project deleted successfully"})
}
