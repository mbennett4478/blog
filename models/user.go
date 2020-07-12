package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName	string
	LastName	string
	Email 		string
	Version		uint		`gorm:"default:0;"`
	BirthDate	*time.Time
	Password	string		`json:"-"`
	Roles		[]Role 		`gorm:"many2many:user_roles;"`
	Posts		[]Post
	Comments	[]Comment
}
