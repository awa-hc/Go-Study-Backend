package handlers

import (
	"net/http"

	"github.com/awa-hc/backend/initializers/database"
	"github.com/awa-hc/backend/initializers/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var user []models.Users
	if err := database.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": user})

}

func GetUser(c *gin.Context) {
	var id = c.Param("id")
	var user models.Users
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
