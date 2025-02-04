package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitnet/internal/pkg/common"
	"gitnet/internal/server/service/git"
)

// route: /api/repo/getFiles [Get]
func GetFiles(ctx *gin.Context) {
	var GetFilesRequest common.GetFilesRequest
	if err := ctx.ShouldBindJSON(&GetFilesRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repoPath := GetFilesRequest.RepoPath
	revision := GetFilesRequest.Revision

	// TODO: userID
	files, err := service.GetFiles("", repoPath, revision)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"files": files})
}
