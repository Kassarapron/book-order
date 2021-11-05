package repository

import (
	"book-order-be/entity"
	"book-order-be/setup"
)

type Book = entity.Book

func InsertBook(book *Book) error {
	if result := setup.Db.Create(book); result.Error != nil {
		return result.Error
	}
	return nil
}

func ListBooks() ([]Book, error) {
	var books []Book
	if result := setup.Db.Preload("BookType").Find(&books); result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func FindBookById(id string) (*Book, error) {
	var book Book
	if result := setup.Db.First(&book, id); result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}
