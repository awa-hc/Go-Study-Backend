package handlers

import (
	"net/http"

	"github.com/awa-hc/backend/initializers/database"
	"github.com/awa-hc/backend/initializers/models"
	"github.com/gin-gonic/gin"
)

func GetProjectComments(c *gin.Context) {
	id := c.Param("id")
	var projectComments []models.ProjectComment

	if err := database.DB.Preload("User").Preload("Project").Preload("Comment").Preload("Comment.User").Where("project_id = ?", id).Find(&projectComments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find project comments"})
		return
	}

	simplifiedProjectComments := make([]gin.H, len(projectComments))
	for i, projectComment := range projectComments {
		simplifiedProjectComments[i] = gin.H{
			"ID":       projectComment.ID,
			"Comment":  projectComment.Comment.Text,
			"ParendID": projectComment.Comment.ParentID,
			"User":     projectComment.Comment.User.Fullname,
		}
	}

	c.JSON(http.StatusOK, gin.H{"projectcomments": simplifiedProjectComments})

	// c.JSON(http.StatusOK, gin.H{"projectcomments": projectComments})
}
