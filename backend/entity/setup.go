package entity
import (
        "gorm.io/gorm"
        "gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {
        return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

  // Migrate the schema
  database.AutoMigrate(
    &Admin{},
    &Company{},
    &BookType{},
    &Book{},
    &BookOrder{},
  )

  db = database

  //Admin Data
  db.Model(&Admin{}).Create(&Admin{
    AdminName : "Chanwit",
    Email: "example01@gmail.com",
  })
  db.Model(&Admin{}).Create(&Admin{
    AdminName : "Kassarapron",
    Email: "example02@gmail.com",
  })
  var chanwit Admin
	var Kassarapron Admin
	db.Raw("SELECT * FROM admins WHERE email = ?", "example01@gmail.com").Scan(&chanwit)
	db.Raw("SELECT * FROM admins WHERE email = ?", "example02@gmail.com").Scan(&Kassarapron)

  
  //Book Data
  db.Model(&Book{}).Create(&Book{
    BookName      : "เจ้าขุนทอง",
    BookNumber    : "AB0152",
    BookPublicher : "สำนักพิมพ์01",
  })
  db.Model(&Book{}).Create(&Book{
    BookName      : "วรรณคดี ม.2",
    BookNumber    : "AH2124",
    BookPublicher : "สำนักพิมพ์02",
  })
  db.Model(&Book{}).Create(&Book{
    BookName      : "แฮรี่พอตเตอร์",
    BookNumber    : "SG2646",
    BookPublicher : "สำนักพิมพ์03",
  })
  db.Model(&Book{}).Create(&Book{
    BookName      : "แฟชั่นเสื้อผ้า",
    BookNumber    : "DK2860",
    BookPublicher : "สำนักพิมพ์04",
    // BookTypeID    : 1,
  })

  // B01 := Book{BookName: "เจ้าขุนทอง",}
  // db.Model(&Book{}).Create(&B01)

  // B02 := Book{BookName: "วรรณคดี ม.2",}
  // db.Model(&Book{}).Create(&B02)

  // B03 := Book{BookName: "แฮรี่พอตเตอร์",}
  // db.Model(&Book{}).Create(&B03)
  
  // B04 := Book{BookName: "แฟชั่นเสื้อผ้า",}
  // db.Model(&Book{}).Create(&B04)

  //BookType Data
  t01 := BookType{TypeName: "หนังสือนิทาน",}
  db.Model(&BookType{}).Create(&t01)
  
  t02 := BookType{TypeName: "หนังสือเรียน",}
  db.Model(&BookType{}).Create(&t02)

  t03 := BookType{TypeName: "นิยาย",}
  db.Model(&BookType{}).Create(&t03)

  t04 := BookType{TypeName: "วรสาร",}
  db.Model(&BookType{}).Create(&t04)

  t05 := BookType{TypeName: "นิตยสาร",}
  db.Model(&BookType{}).Create(&t05)

  t06 := BookType{TypeName: "สารคดี",}
  db.Model(&BookType{}).Create(&t06)

  //Company Data
  Acompany := Company{CompanyName: "ซีเอ็ดบุ๊ค",}
  db.Model(&Company{}).Create(&Acompany)

  Bcompany := Company{CompanyName: "ไทยเสรีการพิมพ์",}
  db.Model(&Company{}).Create(&Bcompany)

  Ccompany := Company{CompanyName: "โรงพิมพ์อักษร",}
  db.Model(&Company{}).Create(&Ccompany)

	// var orderList []*BookOrder
	// db.Model(&BookOrder{}).
	// 	Joins("Admin").
	// 	Joins("Book").
	// 	Joins("Company")

	// for _, ol := range orderList {
	// 	fmt.Printf("Book Order: %v\n", ol.ID)
	// 	fmt.Printf("%v\n", ol.Admin.AdminName)
	// 	fmt.Printf("%v\n", ol.Book.BookName)
	// 	fmt.Printf("%v\n", ol.Company.CompanyName)
	// 	fmt.Println("====")
	// }
}
