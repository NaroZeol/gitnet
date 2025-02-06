package service

import (
	"fmt"

	"gitnet/internal/pkg/common"
	"gitnet/internal/pkg/utils"
	"gitnet/internal/server/model/database"
)

func GetUserList() ([]common.User, error) {
	users, err := model.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(name string, password string) error {
	if _, err := model.GetUserByName(name); err == nil {
		return fmt.Errorf("user %s already exists", name)
	}

	cryptedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	return model.CreateUser(name, cryptedPassword)
}

func DeleteUser(name string) error {
	if _, err := model.GetUserByName(name); err != nil {
		return fmt.Errorf("user %s does not exist", name)
	}

	return model.DeleteUserByName(name)
}
