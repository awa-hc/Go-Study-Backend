package handlers

import (
	"net/http"

	"github.com/awa-hc/backend/initializers/database"
	"github.com/awa-hc/backend/initializers/models"
	"github.com/gin-gonic/gin"
)

func Getuserproject(c *gin.Context) {
	var id = c.Param("id")
	var userproject []models.UserProject

	if err := database.DB.Preload("User").Preload("Project").Where("user_id = ?", id).Find(&userproject).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find userproject"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userproject": userproject})
}
