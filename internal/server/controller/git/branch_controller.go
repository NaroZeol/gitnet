package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitnet/internal/pkg/common"
	"gitnet/internal/server/service/git"
)

// route: /api/repo/getBranches [Get]
func GetBranches(ctx *gin.Context) {
	var GetBranchesRequest common.GetBranchesRequest
	if err := ctx.ShouldBindJSON(&GetBranchesRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: userID
	branches, err := service.GetBranches("", GetBranchesRequest.RepoPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"branches": branches})
}

// route: /api/repo/createBranch [Post]
func CreateBranch(ctx *gin.Context) {
	var CreateBranchRequest common.CreateBranchRequest
	if err := ctx.ShouldBindJSON(&CreateBranchRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: userID
	err := service.CreateBranch("", CreateBranchRequest.RepoPath, CreateBranchRequest.BranchName, CreateBranchRequest.Revision)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// route: /api/repo/deleteBranch [Post]
func DeleteBranch(ctx *gin.Context) {
	var DeleteBranchRequest common.DeleteBranchRequest
	if err := ctx.ShouldBindJSON(&DeleteBranchRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: userID
	err := service.DeleteBranch("", DeleteBranchRequest.RepoPath, DeleteBranchRequest.BranchName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
