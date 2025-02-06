package main

import (
	"github.com/gin-gonic/gin"

	"gitnet/internal/pkg/database"
	"gitnet/internal/server/controller"
)

func main() {
	// Connect to database
	database.Init()

	app := gin.Default()
	// Set Router
	controller.SetRouter(app)

	// Listen and Server in 0.0.0.0:8080
	app.Run(":8080")
}
