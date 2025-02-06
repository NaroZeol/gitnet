package service

import (
	"fmt"
	"strconv"

	"gitnet/internal/pkg/common"
	"gitnet/internal/server/model/database"
	"gitnet/internal/server/model/git"
)

func CreateRepository(userID string, repoPath string) (*common.CreateRepositoryResponse, error) {
	if _, err := model.GetRepositoryByPath(repoPath); err == nil {
		return nil, fmt.Errorf("repository %s already exists", repoPath)
	}

	err := git.CreateRepository(repoPath)
	if err != nil {
		return nil, err
	}

	ownerID, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}

	if err := model.CreateRepository(uint(ownerID), repoPath); err != nil {
		return nil, err
	}

	return &common.CreateRepositoryResponse{Message: fmt.Sprintf("repository %s created", repoPath)}, nil
}

func DeleteRepository(userID string, repoPath string) (*common.DeleteRepositoryResponse, error) {
	isOwner, err := CheckOwnership(userID, repoPath)
	if err != nil {
		return nil, err
	}

	if !isOwner {
		return nil, fmt.Errorf("you are not the owner of repository %s", repoPath)
	}

	err = git.DeleteRepository(repoPath)
	if err != nil {
		return nil, err
	}

	if err := model.DeleteRepositoryByPath(repoPath); err != nil {
		return nil, err
	}

	return &common.DeleteRepositoryResponse{Message: fmt.Sprintf("repository %s deleted", repoPath)}, nil
}
