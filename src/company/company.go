package company

import (
	"github.com/zangeronimo/authenticated/src/db"
)

// GetAll return all register for company
func GetAll() (companies []db.Company, err error) {
	dba := db.Connect()
	defer dba.Close()

	err = dba.Find(&companies).Error

	return
}

// GetOne retorn a specific company
func GetOne(id uint) (company db.Company, err error) {
	dba := db.Connect()
	defer dba.Close()

	err = dba.Find(&company, id).Error

	return
}

// NewCompany receive a company, persist in db and return a company or a error
func NewCompany(company db.Company) (db.Company, error) {
	dba := db.Connect()
	defer dba.Close()

	err := dba.Create(&company).Error
	return company, err
}
