package main

import (
	"github.com/gin-gonic/gin"

	"gitnet/internal/server/controller"
)

func main() {
	app := gin.Default()

	app_api := app.Group("/api")
	app_api_repo := app_api.Group("/repo")
	{
		app_api_repo.GET("/getFiles", controller.GetFiles)
		app_api_repo.GET("/getBranches", controller.GetBranches)
		app_api_repo.POST("/createBranch", controller.CreateBranch)
		app_api_repo.POST("/deleteBranch", controller.DeleteBranch)
	}

	// Listen and Server in 0.0.0.0:8080
	app.Run(":8080")
}
