package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"gitnet/internal/pkg/common"
	"gitnet/internal/pkg/database"
)

func GetUserByName(name string) (common.User, error) {
	var user common.User
	err := database.DB.Where("name = ?", name).First(&user).Error
	return user, err
}

func GetUserByID(id uint) (common.User, error) {
	var user common.User
	err := database.DB.First(&user, id).Error
	return user, err
}

func CreateUser(name, password string) error {
	return database.DB.Create(&common.User{Name: name, Password: password}).Error
}

func UpdateUserPassword(id uint, password string) error {
	return database.DB.Model(&common.User{}).Where("id = ?", id).Update("password", password).Error
}

func DeleteUserByID(id uint) error {
	return database.DB.Delete(&common.User{}, id).Error
}

func DeleteUserByName(name string) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		err := database.DB.Where("name = ?", name).Delete(&common.User{}).Error
		if err != nil {
			return err
		}

		deletedUsername := fmt.Sprintf("deleted_%s_%d", name, time.Now().Unix())
		return database.DB.Unscoped().Model(&common.User{}).Where("name = ?", name).Update("name", deletedUsername).Error
	})
}

func GetAllUsers() ([]common.User, error) {
	var users []common.User
	err := database.DB.Find(&users).Error
	return users, err
}
