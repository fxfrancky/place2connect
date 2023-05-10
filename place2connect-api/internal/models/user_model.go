package models

import (
	"gorm.io/gorm"
)

type User struct {
	LastName      string `gorm:"type:varchar(100);not null"`
	FirstName     string `gorm:"type:varchar(100);not null"`
	Email         string `gorm:"type:varchar(100);uniqueIndex:idx_email;not null"`
	Password      string `gorm:"type:varchar(100);not null"`
	Role          string `gorm:"type:varchar(50);default:'user';not null"`
	Provider      string `gorm:"type:varchar(50);default:'local';not null"`
	Photo         string `gorm:"not null;default:'default.png'"`
	Verified      bool   `gorm:"not null;default:false"`
	IsAdmin       bool   `gorm:"not null;default:false"`
	PicturePath   string `gorm:"type:varchar(255);not null"`
	Friends       []User `gorm:"many2many:user_friends"`
	Location      string `gorm:"type:varchar(255);not null"`
	Occupation    string `gorm:"type:varchar(255);not null"`
	ViewedProfile uint   `gorm:"not null;default:0"`
	Impressions   uint   `gorm:"not null;default:0"`
	gorm.Model
}
