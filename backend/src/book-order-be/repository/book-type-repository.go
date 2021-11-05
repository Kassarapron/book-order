package repository

import (
	"book-order-be/entity"
	"book-order-be/setup"
)

type BookType = entity.BookType

func InsertBookType(bookType *BookType) error {
	if result := setup.Db.Create(bookType); result.Error != nil {
		return result.Error
	}
	return nil
}

func ListBookTypes() ([]BookType, error) {
	var bookTypes []BookType
	if result := setup.Db.Find(&bookTypes); result.Error != nil {
		return nil, result.Error
	}
	return bookTypes, nil
}

// func FindBookTypeById(id string) (*BookType, error) {
// 	var admin BookType
// 	if result := setup.Db.First(&admin, id); result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &admin, nil
// }
