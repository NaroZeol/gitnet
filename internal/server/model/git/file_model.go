package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// GetFiles returns the files in a repository
func GetFiles(repoPath string, revision string) ([]string, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}

	realhash, err := repo.ResolveRevision(plumbing.Revision(revision))
	if err != nil {
		return nil, err
	}

	commit, err := repo.CommitObject(*realhash)
	if err != nil {
		return nil, err
	}

	files, err := commit.Files()
	if err != nil {
		return nil, err
	}
	defer files.Close()

	var fileNames []string
	files.ForEach(func(file *object.File) error {
		fileNames = append(fileNames, file.Name)
		return nil
	})

	return fileNames, nil
}
