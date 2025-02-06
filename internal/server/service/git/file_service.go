package service

import (
	"fmt"
	"gitnet/internal/server/model/git"
)

func GetFiles(userID string, repoPath string, revision string) ([]string, error) {
	isOwner, err := CheckOwnership(userID, repoPath)
	if err != nil {
		return nil, err
	}

	if !isOwner {
		return nil, fmt.Errorf("user %s does not own repository %s", userID, repoPath)
	}

	return git.GetFiles(repoPath, revision)
}
