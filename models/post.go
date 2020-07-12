package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title		string
	Description string 		`gorm:"type:text;"`
	UserID		uint
	Comments	[]Comment
}
