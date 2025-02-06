package service

import (
	"fmt"

	"gitnet/internal/pkg/common"
	"gitnet/internal/pkg/utils"
	"gitnet/internal/server/model/git"
)

func GetBranches(userID string, repoPath string) ([]common.Branch, error) {
	isOwner, err := CheckOwnership(userID, repoPath)
	if err != nil {
		return nil, err
	}

	if !isOwner {
		return nil, fmt.Errorf("user %s does not own repository %s", userID, repoPath)
	}

	return git.GetBranches(repoPath)
}

func CreateBranch(userID string, repoPath string, branchName string, revision string) error {
	isOwner, err := CheckOwnership(userID, repoPath)
	if err != nil {
		return err
	}

	if !isOwner {
		return fmt.Errorf("user %s does not own repository %s", userID, repoPath)
	}

	if !utils.CheckBranchName(branchName) {
		return fmt.Errorf("invalid branch name %s", branchName)
	}
	return git.CreateBranch(repoPath, branchName, revision)
}

func DeleteBranch(userID string, repoPath string, branchName string) error {
	isOwner, err := CheckOwnership(userID, repoPath)
	if err != nil {
		return err
	}

	if !isOwner {
		return fmt.Errorf("user %s does not own repository %s", userID, repoPath)
	}

	if !utils.CheckBranchName(branchName) {
		return fmt.Errorf("invalid branch name %s", branchName)
	}
	return git.DeleteBranch(repoPath, branchName)
}
