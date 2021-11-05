package main

import (
	"book-order-be/handler"
	"book-order-be/setup"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	setup.SetupDB()

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/users", handler.ListUsers)
	r.GET("/users/:name", handler.GetUser)
	r.POST("/users", handler.CreateUser)

	r.POST("/admin", handler.CreateAdmin)     //create new admin
	r.GET("/admin", handler.GetAllAdmins)     //list all admins
	r.GET("/admin/:id", handler.GetAdminById) //get admin by id

	r.POST("/book-order", handler.CreateBookOrder) //create new book-order
	r.GET("/book-order", handler.GetAllBookOrders) //list all  book-orders

	r.POST("/company", handler.CreateCompany)  //create new company
	r.GET("/company", handler.GetAllCompanies) //list all companies

	r.POST("/book-type", handler.CreateBookType) //create new book type
	r.GET("/book-type", handler.GetAllBookTypes) //list all book types

	r.POST("/book", handler.CreateBook) //create new book
	r.GET("/book", handler.GetAllBooks) //create new book

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
