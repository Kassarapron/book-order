package setup

import (
	"book-order-be/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DB() *gorm.DB {
	return Db
}

func SetupDB() {
	_db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db = _db

	if err := Db.AutoMigrate(
		&entity.Admin{},
		&entity.Company{},
		&entity.BookType{},
		&entity.Book{},
		&entity.BookOrder{}); err != nil {
		panic(err)
	}
}
