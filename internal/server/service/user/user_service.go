package service

import (
	"fmt"

	"gitnet/internal/pkg/common"
	"gitnet/internal/pkg/utils"
	"gitnet/internal/server/model/database"
)

func GetUserList() (*common.GetUserListResponse, error) {
	users, err := model.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var response common.GetUserListResponse
	for i := range users {
		response.Users = append(response.Users, users[i].ToResponse())
	}

	return &response, nil
}

func CreateUser(name string, password string) (*common.CreateUserResponse, error) {
	if _, err := model.GetUserByName(name); err == nil {
		return nil, fmt.Errorf("user %s already exists", name)
	}

	cryptedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	if err := model.CreateUser(name, cryptedPassword); err != nil {
		return nil, err
	}

	return &common.CreateUserResponse{Message: fmt.Sprintf("user %s created", name)}, nil
}

func DeleteUser(name string) (*common.DeleteUserResponse, error) {
	if _, err := model.GetUserByName(name); err != nil {
		return nil, fmt.Errorf("user %s does not exist", name)
	}

	if err := model.DeleteUserByName(name); err != nil {
		return nil, err
	}

	return &common.DeleteUserResponse{Message: fmt.Sprintf("user %s deleted", name)}, nil
}
