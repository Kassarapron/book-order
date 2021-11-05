package main
import (
  "github.com/kassarapron/sa-64/controller"
  "github.com/kassarapron/sa-64/entity"
  "github.com/gin-gonic/gin"
)
func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	// Admin Routes
	r.GET("/admin"        , controller.ListAdmins)
	r.GET("/admin/:id"    , controller.GetAdmin)
	r.POST("/admin"       , controller.CreateAdmin)
	r.PATCH("/admin"      , controller.UpdateAdmin)
	r.DELETE("/admin/:id" , controller.DeleteAdmin)
   // Book Routes
   r.GET("/book"        , controller.ListBooks)
   r.GET("/book/:id"    , controller.GetBook)
   r.POST("/book"       , controller.CreateBook)
   r.PATCH("/book"      , controller.UpdateBook)
   r.DELETE("/book/:id" , controller.DeleteBook)
   // BookType Routes
   r.GET("/booktype"        , controller.ListBookTypes)
   r.GET("/booktype/:id"    , controller.GetBookType)
   r.POST("/booktype"       , controller.CreateBookType)
   r.PATCH("/booktype"      , controller.UpdateBookType)
   r.DELETE("/booktype/:id" , controller.DeleteBookType)
   // Company Routes
   r.GET("/company"        , controller.ListCompanys)
   r.GET("/company/:id"    , controller.GetCompany)
   r.POST("/company"       , controller.CreateCompany)
   r.PATCH("/company"      , controller.UpdateCompany)
   r.DELETE("/company/:id" , controller.DeleteCompany)
	// BookOrder Routes
	r.GET("/bookorder"        , controller.ListBookOrders)
	r.GET("/bookorder/:id"    , controller.GetBookOrder)
	r.POST("/bookorder"       , controller.CreateBookOrder)
	r.PATCH("/bookorder"      , controller.UpdateBookOrder)
	r.DELETE("/bookorder/:id" , controller.DeleteBookOrder)
  // Run the server
	r.Run()
  }
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