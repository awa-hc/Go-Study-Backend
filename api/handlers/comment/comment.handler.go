package handlers

import (
	"fmt"
	"net/http"

	"github.com/awa-hc/backend/initializers/database"
	"github.com/awa-hc/backend/initializers/models"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	user, _ := c.Get("user")
	userid := user.(models.Users).ID

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	comment.CreatedByID = userid

	if comment.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment can not be empty"})
		return
	}

	if comment.CreatedByID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user creator is not valid"})
		return
	}

	if comment.ProjectID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "project not valid"})
		return
	}

	if comment.ParentID != nil {
		var parentComment models.Comment
		if err := database.DB.Where("id = ?", comment.ParentID).First(&parentComment).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "parent comment not valid"})
			return
		}
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create the comment"})
		return
	}

	var projectComment models.ProjectComment
	projectComment.CommentID = comment.ID
	projectComment.ProjectID = comment.ProjectID
	projectComment.UserID = comment.CreatedByID

	if err := database.DB.Create(&projectComment).Error; err != nil {
		database.DB.Delete(&comment)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create the project comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment created successfully"})

}

func GetComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	var replies []models.Comment

	if err := database.DB.Preload("User").Preload("Project").Where("id = ?", id).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find comment"})
		return
	}
	if err := database.DB.Preload("User").Preload("Project").Where("parent_id = ?", id).Find(&replies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find replies"})
		return
	}

	comment.ChildComments = replies

	simplifiedReplies := make([]gin.H, len(replies))
	for i, reply := range replies {
		simplifiedReplies[i] = gin.H{
			"Text":     reply.Text,
			"ParendID": reply.ParentID,
			"User":     reply.User.Fullname,
		}
	}

	c.JSON(http.StatusOK, gin.H{"comment": simplifiedReplies})
}

func UpdateComment(c *gin.Context) {
	var id = c.Param("id")
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}
	if comment.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment can not be empty"})
		return
	}

	if err := database.DB.Model(&comment).Where("id = ?", id).Updates(comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment updated successfully"})

}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	var projectComment models.ProjectComment

	if err := database.DB.Where("id = ?", id).Delete(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not delete comment"})
		return
	}

	if err := database.DB.Where("comment_id = ?", id).Delete(&projectComment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not delete project comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "comment deleted successfully"})
}

func GetCommentReply(c *gin.Context) {
	var commentid = c.Param("id")
	var replyid = c.Param("ids")
	fmt.Println(replyid)
	var comment models.Comment
	var replies []models.Comment

	if err := database.DB.Preload("User").Preload("Project").Where("parent_id = ? AND id = ?", commentid, replyid).Find(&replies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error in finding replies"})
		return
	}
	comment.ChildComments = replies

	if len(replies) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find replies"})
		return
	}

	simplifiedReplies := make([]gin.H, len(replies))
	for i, reply := range replies {
		simplifiedReplies[i] = gin.H{
			"Text":     reply.Text,
			"ParendID": reply.ParentID,
			"User":     reply.User.Fullname,
		}
	}

	c.JSON(http.StatusOK, gin.H{"reply": simplifiedReplies})

}

func GetCommentsReplies(c *gin.Context) {
	var id = c.Param("id")
	var comment models.Comment
	var replies []models.Comment

	if err := database.DB.Preload("User").Preload("Project").Where("id = ?", id).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find comment"})
		return
	}

	if err := database.DB.Preload("User").Preload("Project").Where("parent_id = ?", id).Find(&replies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not find replies"})
		return
	}

	comment.ChildComments = replies

	simplifiedReplies := make([]gin.H, len(replies))
	for i, reply := range replies {
		simplifiedReplies[i] = gin.H{
			"Text":     reply.Text,
			"ParendID": reply.ParentID,
			"User":     reply.User.Fullname,
		}
	}

	c.JSON(http.StatusOK, gin.H{"replies": simplifiedReplies})

}
