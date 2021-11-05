package repository

import (
	"book-order-be/entity"
	"book-order-be/setup"
)

type BookOrder = entity.BookOrder

func InsertBookOrder(bookOrder *BookOrder) error {
	if result := setup.Db.Create(bookOrder); result.Error != nil {
		return result.Error
	}
	return nil
}

func ListBookOrders() ([]BookOrder, error) {
	var bookOrders []BookOrder
	if result := setup.Db.Preload("Admin").Preload("Company").Preload("Book").Preload("Book.BookType").Find(&bookOrders); result.Error != nil {
		return nil, result.Error
	}
	return bookOrders, nil
}

// func FindBookOrderById(id string) (*BookOrder, error) {
// 	var bookOrder BookOrder
// 	if result := setup.Db.First(&bookOrder, id); result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &bookOrder, nil
// }
