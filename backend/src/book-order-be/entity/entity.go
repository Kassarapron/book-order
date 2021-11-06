package entity

import (
	"gorm.io/gorm"
)

//del this
type Admin struct {
	gorm.Model
	AdminName  string
	Email      string      `gorm:"uniqeIndex"`
	BookOrders []BookOrder `gorm:"foreignKey:AdminID`
}
type Company struct {
	gorm.Model
	CompanyName string
	BookOrders  []BookOrder `gorm:"foreignKey:CompanyID`
}
type BookType struct {
	gorm.Model
	TypeName string
	Book     []Book `gorm:"foreignKey:BookTypeID`
}
type Book struct {
	gorm.Model
	BookName      string
	BookNumber    string `gorm:"uniqeIndex"`
	BookPublicher string
	BookOrders    []BookOrder `gorm:"foreignKey:BookID`
	//BookTypeID ทำหน้าที่เป็น FK
	BookTypeID *uint
	BookType   BookType `gorm:references:id`
}
type BookOrder struct {
	gorm.Model
	Quantity uint
	//AdminID ทำหน้าที่เป็น FK
	AdminID *uint
	Admin   Admin `gorm:references:id`
	//CompanyID ทำหน้าที่เป็น FK
	CompanyID *uint
	Company   Company `gorm:references:id`
	//BookID ทำหน้าที่เป็น FK
	BookID *uint
	Book   Book `gorm:references:id`
}
