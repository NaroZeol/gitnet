package model

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"gitnet/internal/pkg/common"
)

func GetBranches(repoPath string) ([]common.Branch, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}

	branches, err := repo.Branches()
	if err != nil {
		return nil, err
	}
	defer branches.Close()

	var branchList []common.Branch
	branches.ForEach(func(ref *plumbing.Reference) error {
		branchList = append(branchList, common.Branch{
			Name: ref.Name().Short(),
			Hash: ref.Hash().String(),
		})
		return nil
	})

	return branchList, nil
}

func CreateBranch(repoPath string, branchName string, revision string) error {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	// check if branch already exists
	_, err = repo.Storer.Reference(plumbing.ReferenceName("refs/heads/" + branchName))
	if err != plumbing.ErrReferenceNotFound {
		return fmt.Errorf("branch %s already exists", branchName)
	}

	hash, err := repo.ResolveRevision(plumbing.Revision(revision))
	if err != nil {
		return err
	}

	ref := plumbing.NewHashReference(plumbing.ReferenceName("refs/heads/"+branchName), *hash)

	return repo.Storer.SetReference(ref)
}

func DeleteBranch(repoPath string, branchName string) error {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	// check if branch exists
	_, err = repo.Storer.Reference(plumbing.ReferenceName("refs/heads/" + branchName))
	if err == plumbing.ErrReferenceNotFound {
		return fmt.Errorf("branch %s does not exist", branchName)
	}

	// check if branch is current HEAD
	headRef, err := repo.Head()
	if err != nil {
		return err
	}
	if headRef.Name().Short() == branchName {
		return fmt.Errorf("branch %s is current HEAD", branchName)
	}

	return repo.Storer.RemoveReference(plumbing.ReferenceName("refs/heads/" + branchName))
}
