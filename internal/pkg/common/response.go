package common

// file_controller.go

type GetRepoFilesResponse struct {
	Files []string `json:"files"`
}

// branch_controller.go

type GetRepoBranchesResponse struct {
	Branches []Branch `json:"branches"`
}

type CreateBranchResponse struct {
	Message string `json:"message"`
}

type DeleteBranchResponse struct {
	Message string `json:"message"`
}

// repo_controller.go

type Repository struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	OwnnerID uint   `json:"ownner_id"`
}

type CreateRepositoryResponse struct {
	Message string `json:"message"`
}

type DeleteRepositoryResponse struct {
	Message string `json:"message"`
}

// user_controller.go

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetUserListResponse struct {
	Users []User `json:"users"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
