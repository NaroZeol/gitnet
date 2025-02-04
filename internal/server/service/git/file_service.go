package service

import (
	"gitnet/internal/server/model/git"
)

func GetFiles(userID string, repoPath string, revision string) ([]string, error) {
	// TODO: Check if user has access to the repository
	return model.GetFiles(repoPath, revision)
}
