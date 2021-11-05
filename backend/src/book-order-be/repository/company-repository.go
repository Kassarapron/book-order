package repository

import (
	"book-order-be/entity"
	"book-order-be/setup"
)

type Company = entity.Company

func InsertCompany(company *Company) error {
	if result := setup.Db.Create(company); result.Error != nil {
		return result.Error
	}
	return nil
}

func ListCompanies() ([]Company, error) {
	var companies []Company
	if result := setup.Db.Find(&companies); result.Error != nil {
		return nil, result.Error
	}
	return companies, nil
}

func FindCompanyById(id string) (*Company, error) {
	var company Company
	if result := setup.Db.First(&company, id); result.Error != nil {
		return nil, result.Error
	}
	return &company, nil
}
