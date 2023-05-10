package models

import "gorm.io/gorm"

type Post struct {
	User            User `gorm:"foreignKey:UserID"`
	UserID          uint
	FirstName       string `gorm:"type:varchar(100);not null"`
	LastName        string `gorm:"type:varchar(100);not null"`
	Location        string `gorm:"type:varchar(100);not null"`
	Description     string `gorm:"type:varchar(255);not null"`
	PicturePath     string `gorm:"type:varchar(255);"`
	UserPicturePath string `gorm:"type:varchar(255);"`
	Likes           []Like `gorm:"constraint:OnDelete:CASCADE;"`
	Comments        []Comment
	gorm.Model
}
