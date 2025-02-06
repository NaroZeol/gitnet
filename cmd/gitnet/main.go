package main

import (
	"github.com/gin-gonic/gin"

	"gitnet/internal/pkg/config"
	"gitnet/internal/pkg/database"
	"gitnet/internal/server/controller"
)

func main() {
	// Load config
	conf := config.LoadConfig("config.json")

	// Connect to database
	database.Init(conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.DBName)

	app := gin.Default()
	// Set Router
	controller.SetRouter(app)

	// Listen and Server in 0.0.0.0:8080
	app.Run(":" + conf.ServerPort)
}
