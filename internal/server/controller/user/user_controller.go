package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitnet/internal/pkg/common"
	"gitnet/internal/server/service/user"
)

func GetUserList(c *gin.Context) {
	resp, err := service.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func CreateUser(c *gin.Context) {
	var req common.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := service.CreateUser(req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context) {
	var req common.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reps, err := service.DeleteUser(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reps)
}
