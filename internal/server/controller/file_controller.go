package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitnet/internal/server/service"
)

func GetRepoFiles(ctx *gin.Context) {
	repoPath := ctx.Param("repoPath")
	revision := ctx.Param("revision")

	files, err := service.GetRepoFiles(repoPath, revision)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"files": files})
}
