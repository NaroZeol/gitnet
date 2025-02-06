package controller

import (
	"github.com/gin-gonic/gin"

	git "gitnet/internal/server/controller/git"
	user "gitnet/internal/server/controller/user"
)

func SetRouter(app *gin.Engine) {
	app_api := app.Group("/api")

	app_api_repo := app_api.Group("/repo")
	{
		app_api_repo.POST("/createRepository", git.CreateRepository)
		app_api_repo.POST("/deleteRepository", git.DeleteRepository)
	}

	app_api_repo_file := app_api_repo.Group("/file") // /api/repo/file
	{
		app_api_repo_file.GET("/getFiles", git.GetFiles)
	}

	app_api_repo_branch := app_api_repo.Group("/branch") // /api/repo/branch
	{
		app_api_repo_branch.GET("/getBranches", git.GetBranches)
		app_api_repo_branch.POST("/createBranch", git.CreateBranch)
		app_api_repo_branch.POST("/deleteBranch", git.DeleteBranch)
	}

	app_api_user := app_api.Group("/user")
	{
		app_api_user.GET("getUserList", user.GetUserList)
		app_api_user.POST("createUser", user.CreateUser)
		app_api_user.POST("deleteUser", user.DeleteUser)
	}
}
