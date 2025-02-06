package service

import (
	"gitnet/internal/server/model/database"
	"strconv"
)

func CheckOwnership(userID string, repoPath string) (bool, error) {
	repo, err := model.GetRepositoryByPath(repoPath)
	if err != nil {
		return false, err
	}

	ownerID, err := strconv.Atoi(userID)
	if err != nil {
		return false, err
	}

	return repo.OwnerID == uint(ownerID), nil
}
