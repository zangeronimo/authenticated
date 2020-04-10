package models

import (
	"github.com/jinzhu/gorm"
)

// Company struct
type Company struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)"`
	Products []Product
	Users    []User
}

// Product struct
type Product struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100)"`
	CompanyID uint
}

// User struct
type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"type:varchar(100);unique_index"`
	CompanyID uint
	Products  []Product `gorm:"many2many:user_product;"`
}
