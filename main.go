package main

import (
	"os"
	"os/exec"

	auth "github.com/awa-hc/backend/api/auth"
	commenthandler "github.com/awa-hc/backend/api/handlers/comment"
	projecthandler "github.com/awa-hc/backend/api/handlers/project"
	projectcommenthandler "github.com/awa-hc/backend/api/handlers/projectcomment"
	sendemail "github.com/awa-hc/backend/api/handlers/sendemail"
	taskhandler "github.com/awa-hc/backend/api/handlers/task"
	taskprojecthandler "github.com/awa-hc/backend/api/handlers/taskproject"
	userhandler "github.com/awa-hc/backend/api/handlers/user"
	uprojecthandler "github.com/awa-hc/backend/api/handlers/userproject"

	"github.com/awa-hc/backend/api/middleware"
	_ "github.com/awa-hc/backend/docs"
	"github.com/awa-hc/backend/initializers"
	"github.com/awa-hc/backend/initializers/database"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDB()
	clearscreen()

}

func clearscreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// @title tag Service Api
// @version 1.0
// @description service api for gostudy using gin
// @host localhost:8080
// @BasePath /
func main() {
	var port = os.Getenv("PORT")

	route := gin.Default()

	// Add swagger
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authGroup := route.Group("/auth")

	authGroup.POST("/signup", auth.SignUp)
	authGroup.POST("/login", auth.Login)
	authGroup.GET("/validate", middleware.RequireAuth, auth.Validate)

	userGroup := route.Group("/user")
	{
		userGroup.GET("/", userhandler.GetUsers)
		userGroup.GET("/:id", userhandler.GetUser)
	}

	projectGroup := route.Group("/project")
	projectGroup.Use(middleware.RequireAuth)
	{
		projectGroup.POST("/", projecthandler.CreateProject)
		projectGroup.GET("/", projecthandler.GetProjects)
		projectGroup.GET("/:id", projecthandler.GetProject)
		projectGroup.DELETE("/:id", projecthandler.DeleteProject)
	}
	userprojecetGroup := route.Group("/userproject")
	userprojecetGroup.Use(middleware.RequireAuth)
	{
		userprojecetGroup.GET("/:id", uprojecthandler.Getuserproject)
	}

	taskGroup := route.Group("/task")
	taskGroup.Use(middleware.RequireAuth)
	{
		taskGroup.POST("/", taskhandler.CreateTask)
		taskGroup.GET("/", taskhandler.GetTasks)
		taskGroup.GET("/:id", taskhandler.GetTask)
	}

	taskprojectGroup := route.Group("/taskproject")
	taskprojectGroup.Use(middleware.RequireAuth)
	{
		taskprojectGroup.GET("/:id", taskprojecthandler.GetTaskProject)

	}

	commentGroup := route.Group("/comment")
	commentGroup.Use(middleware.RequireAuth)
	{
		commentGroup.POST("/", commenthandler.CreateComment)
		commentGroup.GET("/:id", commenthandler.GetComment)
		commentGroup.GET("/:id/reply/:ids", commenthandler.GetCommentReply)
		commentGroup.GET("/:id/replies", commenthandler.GetCommentsReplies)
		commentGroup.PUT("/:id", commenthandler.UpdateComment)
		commentGroup.DELETE("/:id", commenthandler.DeleteComment)
	}

	projectCommentGroup := route.Group("/projectcomment")
	projectCommentGroup.Use(middleware.RequireAuth)
	{
		projectCommentGroup.GET("/:id", projectcommenthandler.GetProjectComments)
	}

	route.GET("/send-email", sendemail.SendEmail)

	route.Run(":" + port)

}
