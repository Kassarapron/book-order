package repository

import (
	"book-order-be/entity"
	"book-order-be/setup"
)

type Admin = entity.Admin

func InsertAdmin(admin *Admin) error {
	if result := setup.Db.Create(admin); result.Error != nil {
		return result.Error
	}
	return nil
}

func ListAdmins() ([]Admin, error) {
	var admins []Admin
	if result := setup.Db.Find(&admins); result.Error != nil {
		return nil, result.Error
	}
	return admins, nil
}

func FindAdminById(id string) (*Admin, error) {
	var admin Admin
	if result := setup.Db.First(&admin, id); result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}
