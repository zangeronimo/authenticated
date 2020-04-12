package graphql

import (
	"github.com/zangeronimo/authenticated/db"
	"github.com/zangeronimo/authenticated/company"
)

func getAllCompanies() ([]db.Company, error) {
	return company.GetAll(), nil
}

func getOneCompany(id uint) (db.Company, error) {
	return company.GetOne(id), nil
}

func addCompany(name string) (db.Company, error) {
	return company.NewCompany(name), nil
}