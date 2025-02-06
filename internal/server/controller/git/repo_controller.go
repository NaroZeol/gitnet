package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitnet/internal/pkg/common"
	"gitnet/internal/server/service/git"
)

// route: /api/repo/createRepositoy [Get]
func CreateRepository(ctx *gin.Context) {
	var CreateRepositoryRequest common.CreateRepositoryRequest
	if err := ctx.ShouldBindJSON(&CreateRepositoryRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: get user id from token
	response, err := service.CreateRepository("30", CreateRepositoryRequest.RepoPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// route: /api/repo/deleteRepository [Get]
func DeleteRepository(ctx *gin.Context) {
	var DeleteRepositoryRequest common.DeleteRepositoryRequest
	if err := ctx.ShouldBindJSON(&DeleteRepositoryRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: get user id from token
	response, err := service.DeleteRepository("30", DeleteRepositoryRequest.RepoPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
