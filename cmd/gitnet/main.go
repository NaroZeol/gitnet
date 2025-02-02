package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	// Listen and Server in 0.0.0.0:8080
	app.Run(":8080")
}
