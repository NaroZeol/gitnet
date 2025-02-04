package service

import (
	"fmt"

	"gitnet/internal/pkg/common"
	"gitnet/internal/pkg/utils"
	"gitnet/internal/server/model/git"
)

func GetBranches(userID string, repoPath string) ([]common.Branch, error) {
	// TODO: Check if user has access to the repository

	return model.GetBranches(repoPath)
}

func CreateBranch(userID string, repoPath string, branchName string, revision string) error {
	// TODO: Check if user has access to the repository

	if !utils.CheckBranchName(branchName) {
		return fmt.Errorf("invalid branch name %s", branchName)
	}
	return model.CreateBranch(repoPath, branchName, revision)
}

func DeleteBranch(userID string, repoPath string, branchName string) error {
	// TODO: Check if user has access to the repository

	if !utils.CheckBranchName(branchName) {
		return fmt.Errorf("invalid branch name %s", branchName)
	}
	return model.DeleteBranch(repoPath, branchName)
}
