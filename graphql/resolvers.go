package graphql

import (
	"github.com/zangeronimo/authenticated/company"
	"github.com/zangeronimo/authenticated/db"
)

func getAllCompanies() ([]db.Company, error) {
	return company.GetAll(), nil
}

func getOneCompany(id uint) (db.Company, error) {
	return company.GetOne(id), nil
}
