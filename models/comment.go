package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Description 	string
	UserID			uint
	PostID			uint
}
