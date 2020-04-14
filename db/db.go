package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Company struct
type Company struct {
	ID       uint   ` json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name     string ` json:"name" gorm:"type:varchar(100)"`
	Email    string ` json:"email" gorm:"type:varchar(100);unique_index;not null"`
	Phone    string ` json:"phone" gorm:"type:varchar(30)"`
	Products []Product
	Users    []User
	gorm.Model
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

// New start a connection on a dabasase
func New() {
	db = Connect()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Company{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
}

// Connect return a pointer to MySQL Connection
func Connect() *gorm.DB {
	dbBase := os.Getenv("DB_BASE")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	db, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbBase))
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
