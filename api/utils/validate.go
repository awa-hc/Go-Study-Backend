package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ValidateRequiredStringFields(c *gin.Context, field, value string) bool {
	if value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": field + " is required"})
		return false
	}
	return true
}

func ValidateRequiredTimeFields(c *gin.Context, field string, value time.Time) bool {
	if value.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": field + " is required"})
		return false
	}
	return true
}

func ValidateTimeFields(c *gin.Context, field1 time.Time, field2 time.Time, field3 time.Time) bool {

	if (field2).Unix() < (field1).Unix() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "test date must be after start date"})
		return false
	}

	if (field3).Unix() < (field2).Unix() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end date must be after test date"})
		return false
	}

	if (field3).Unix() < (field1).Unix() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end date must be after start date"})
		return false
	}
	return true
}
