package database

import (
	"time"

	"gitnet/internal/pkg/common"
	"gorm.io/gorm"
)

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type User struct {
	gorm.Model `json:"-"`
	ID         uint           `gorm:"primary_key:id"`
	Name       string         `gorm:"column:name"`
	Password   string         `gorm:"column:password"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) ToResponse() common.User {
	return common.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

type Repository struct {
	gorm.Model `json:"-"`
	ID         uint           `gorm:"primary_key:id"`
	Name       string         `gorm:"column:name"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`

	OwnnerID uint `gorm:"foreignkey:UserID" json:"ownner_id"`
}

func (Repository) TableName() string {
	return "repository"
}
