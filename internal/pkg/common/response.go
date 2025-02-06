package common

// file_controller.go

type GetRepoFilesResponse struct {
	Files []string `json:"files"`
}

// repo_controller.go

type GetRepoBranchesResponse struct {
	Branches []Branch `json:"branches"`
}

type CreateBranchResponse struct {
	Message string `json:"message"`
}

type DeleteBranchResponse struct {
	Message string `json:"message"`
}

// user_controller.go

type GetUserListResponse struct {
	Users []User `json:"users"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
