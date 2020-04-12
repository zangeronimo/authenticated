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

func NewCompany(name string) db.Company {
	dba := db.Connect()
	defer dba.Close()

	var company db.Company
	company.Name = name
	dba.Create(&company)

	return company
}
