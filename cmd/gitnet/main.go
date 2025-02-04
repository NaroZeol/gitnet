package main

import (
	"github.com/gin-gonic/gin"

	"gitnet/internal/server/controller/git"
)

func main() {
	app := gin.Default()

	app_api := app.Group("/api")
	app_api_repo := app_api.Group("/repo")

	app_api_repo_file := app_api_repo.Group("/file") // /api/repo/file
	{
		app_api_repo_file.GET("/getFiles", controller.GetFiles)
	}

	app_api_repo_branch := app_api_repo.Group("/branch") // /api/repo/branch
	{
		app_api_repo_branch.GET("/getBranches", controller.GetBranches)
		app_api_repo_branch.POST("/createBranch", controller.CreateBranch)
		app_api_repo_branch.POST("/deleteBranch", controller.DeleteBranch)
	}

	// Listen and Server in 0.0.0.0:8080
	app.Run(":8080")
}
