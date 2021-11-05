package repository

import (
	"book-order-be/entity"
	"book-order-be/setup"
)

type User = entity.User

func InsertUser(user *User) error {
	if result := setup.Db.Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func FindUsers(age uint) ([]User, error) {
	var users []User
	if result := setup.Db.Where("Age > ?", age).Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func FindByName(name string) (*User, error) {
	var user User
	if result := setup.Db.Where(&User{Name: name}).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
