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

	//Admin Data
	Db.Model(&entity.Admin{}).Create(&entity.Admin{
		AdminName: "Chanwit",
		Email:     "example01@gmail.com",
	})
	Db.Model(&entity.Admin{}).Create(&entity.Admin{
		AdminName: "Kassarapron",
		Email:     "example02@gmail.com",
	})

	//BookType Data
	t01 := entity.BookType{
		TypeName: "นิทาน",
	}
	Db.Model(&entity.BookType{}).Create(&t01)
	t02 := entity.BookType{
		TypeName: "หนังสือเรียน",
	}
	Db.Model(&entity.BookType{}).Create(&t02)
	t03 := entity.BookType{
		TypeName: "นิยาย",
	}
	Db.Model(&entity.BookType{}).Create(&t03)
	t05 := entity.BookType{
		TypeName: "นิตยสาร",
	}
	Db.Model(&entity.BookType{}).Create(&t05)

	//Company Data
	Acompany := entity.Company{CompanyName: "ซีเอ็ดบุ๊ค"}
	Db.Model(&entity.Company{}).Create(&Acompany)

	Bcompany := entity.Company{CompanyName: "ไทยเสรีการพิมพ์"}
	Db.Model(&entity.Company{}).Create(&Bcompany)

	Ccompany := entity.Company{CompanyName: "โรงพิมพ์อักษร"}
	Db.Model(&entity.Company{}).Create(&Ccompany)

	//Book Data
	Db.Model(&entity.Book{}).Create(&entity.Book{
		BookName:      "เจ้าขุนทอง",
		BookNumber:    "AB0152",
		BookPublicher: "สำนักพิมพ์01",
		BookAuthor:    "คนแต่งเจ้าขุนทอง",
		BookTypeID:    &t01.ID,
	})
	Db.Model(&entity.Book{}).Create(&entity.Book{
		BookName:      "วรรณคดี ม.2",
		BookNumber:    "AH2124",
		BookPublicher: "สำนักพิมพ์02",
		BookAuthor:    "Harry Potter",
		BookTypeID:    &t02.ID,
	})
	Db.Model(&entity.Book{}).Create(&entity.Book{
		BookName:      "แฮรี่พอตเตอร์",
		BookNumber:    "SG2646",
		BookPublicher: "สำนักพิมพ์03",
		BookAuthor:    "The duckky",
		BookTypeID:    &t03.ID,
	})
	Db.Model(&entity.Book{}).Create(&entity.Book{
		BookName:      "แฟชั่นเสื้อผ้า",
		BookNumber:    "DK2860",
		BookPublicher: "สำนักพิมพ์04",
		BookAuthor:    "Eye",
		BookTypeID:    &t05.ID,
	})

}
