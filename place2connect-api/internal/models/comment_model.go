package models

import "gorm.io/gorm"

type Comment struct {
	User               User `gorm:"foreignKey:UserID"`
	UserID             uint
	Post               Post `gorm:"foreignKey:PostID"`
	PostID             uint
	CommentDescription string
	gorm.Model
}
