package models

import "gorm.io/gorm"

type Like struct {
	User          User `gorm:"foreignKey:UserID"`
	UserID        uint
	Post          Post `gorm:"foreignKey:PostID"`
	PostID        uint
	IsLikablePost bool
	gorm.Model
}
