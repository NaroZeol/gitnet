package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"gitnet/internal/pkg/database"
)

func GetRepositoryByName(name string) (database.Repository, error) {
	var repository database.Repository
	err := database.DB.Where("name = ?", name).First(&repository).Error
	return repository, err
}

func GetRepositoryByID(id uint) (database.Repository, error) {
	var repository database.Repository
	err := database.DB.First(&repository, id).Error
	return repository, err
}

func GetRepositoryByPath(path string) (database.Repository, error) {
	var repository database.Repository
	err := database.DB.Where("path = ?", path).First(&repository).Error
	return repository, err
}

func GetRepositoryByOwnerID(ownerID uint) ([]database.Repository, error) {
	var repositories []database.Repository
	err := database.DB.Where("owner_id = ?", ownerID).Find(&repositories).Error
	return repositories, err
}

func CreateRepository(ownerID uint, path string) error {
	return database.DB.Create(&database.Repository{Path: path, OwnerID: ownerID}).Error
}

func UpdateRepositoryName(id uint, name string) error {
	return database.DB.Model(&database.Repository{}).Where("id = ?", id).Update("name", name).Error
}

func UpdateRepositoryPath(id uint, path string) error {
	return database.DB.Model(&database.Repository{}).Where("id = ?", id).Update("path", path).Error
}

func DeleteRepositoryByPath(path string) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		err := database.DB.Where("path = ?", path).Delete(&database.Repository{}).Error
		if err != nil {
			return err
		}

		deletedRepoPath := fmt.Sprintf("deleted_%s_%d", path, time.Now().Unix())
		return database.DB.Unscoped().Model(&database.Repository{}).Where("path = ?", path).Update("path", deletedRepoPath).Error
	})
}
