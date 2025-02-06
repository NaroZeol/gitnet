package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitnet/internal/pkg/common"
	"gitnet/internal/server/service/user"
)

func GetUserList(c *gin.Context) {
	users, err := service.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp common.GetUserListResponse
	resp.Users = users
	c.JSON(http.StatusOK, resp)
}

func CreateUser(c *gin.Context) {
	var req common.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateUser(req.Name, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp common.CreateUserResponse
	resp.Message = fmt.Sprintf("user %s created", req.Name)
	c.JSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context) {
	var req common.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.DeleteUser(req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp common.DeleteUserResponse
	resp.Message = fmt.Sprintf("user %s deleted", req.Name)
	c.JSON(http.StatusOK, resp)
}
