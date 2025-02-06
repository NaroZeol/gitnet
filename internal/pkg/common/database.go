package common

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID         uint           `gorm:"primary_key:id" json:"id"`
	Name       string         `gorm:"column:name" json:"name"`
	Password   string         `gorm:"column:password" json:"-"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (User) TableName() string {
	return "user"
}

type Repository struct {
	gorm.Model `json:"-"`
	ID         uint           `gorm:"primary_key:id" json:"id"`
	Name       string         `gorm:"column:name" json:"name"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`

	OwnnerID uint `gorm:"foreignkey:UserID" json:"ownner_id"`
}

func (Repository) TableName() string {
	return "repository"
}
