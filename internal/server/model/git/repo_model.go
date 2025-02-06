package git

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func CreateRepository(path string) error {
	// Create a new repository
	_, err := git.PlainInit(path, true)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRepository(path string) error {
	// Delete the repository
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}
