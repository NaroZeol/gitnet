package common

// file_controller.go

type GetFilesRequest struct {
	RepoPath string `json:"repoPath"`
	Revision string `json:"revision"`
}

// branch_controller.go

type GetBranchesRequest struct {
	RepoPath string `json:"repoPath"`
}

type CreateBranchRequest struct {
	RepoPath   string `json:"repoPath"`
	BranchName string `json:"branchName"`
	Revision   string `json:"revision"`
}

type DeleteBranchRequest struct {
	RepoPath   string `json:"repoPath"`
	BranchName string `json:"branchName"`
}

// repo_controller.go

type CreateRepositoryRequest struct {
	RepoPath string `json:"repoPath"`
}

type DeleteRepositoryRequest struct {
	RepoPath string `json:"repoPath"`
}

// user_controller.go

type CreateUserRequest struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

type DeleteUserRequest struct {
	Name string `json:"username"`
}
