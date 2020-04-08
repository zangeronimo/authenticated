package models

import (
	"github.com/jinzhu/gorm"
)

// User struct for Users
type User struct {
	gorm.Model
	Code  string
	Price uint
	Email string `gorm:"type:varchar(100);unique_index"`
}
