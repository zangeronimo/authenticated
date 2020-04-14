package company

import (
	"github.com/zangeronimo/authenticated/db"
)

func GetAll() []db.Company {
	dba := db.Connect()
	defer dba.Close()

	var companies []db.Company
	dba.Find(&companies)

	return companies
}

func GetOne(id uint) db.Company {
	dba := db.Connect()
	defer dba.Close()

	var company db.Company
	dba.Find(&company, id)

	return company
}

// NewCompany receive a company, persist in db and return a company or a error
func NewCompany(company db.Company) (db.Company, error) {
	dba := db.Connect()
	defer dba.Close()

	err := dba.Create(&company).Error
	return company, err
}
