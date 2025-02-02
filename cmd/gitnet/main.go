package main

import (
	"github.com/gin-gonic/gin"

	"gitnet/internal/server/controller"
)

func main() {
	app := gin.Default()

	app.GET("/repos/:repoPath/revision/:revision/files", controller.GetRepoFiles)

	// Listen and Server in 0.0.0.0:8080
	app.Run(":8080")
}
