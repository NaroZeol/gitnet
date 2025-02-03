package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitnet/internal/pkg/common"
	"gitnet/internal/server/service"
)

// route: GET /api/repos/getBranches [Get]
func GetBranches(ctx *gin.Context) {
	var GetBranchesRequest common.GetBranchesRequest
	if err := ctx.ShouldBindJSON(&GetBranchesRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branches, err := service.GetBranches(GetBranchesRequest.RepoPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"branches": branches})
}

// route: GET /api/repos/createBranch [Post]
func CreateBranch(ctx *gin.Context) {
	var CreateBranchRequest common.CreateBranchRequest
	if err := ctx.ShouldBindJSON(&CreateBranchRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateBranch(CreateBranchRequest.RepoPath, CreateBranchRequest.BranchName, CreateBranchRequest.Revision)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func DeleteBranch(ctx *gin.Context) {
	var DeleteBranchRequest common.DeleteBranchRequest
	if err := ctx.ShouldBindJSON(&DeleteBranchRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.DeleteBranch(DeleteBranchRequest.RepoPath, DeleteBranchRequest.BranchName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
