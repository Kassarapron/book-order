package controller
import (
        "github.com/kassarapron/sa-64/entity"
        "github.com/gin-gonic/gin"
        "net/http"
)
// POST /books
func CreateBook(c *gin.Context) {
	var book entity.Book
	if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if err := entity.DB().Create(&book).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}
// GET /book/:id
func GetBook(c *gin.Context) {
	var book entity.Book
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM books WHERE id = ?", id).Scan(&book).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}
// GET /books
func ListBooks(c *gin.Context) {
	var books []entity.Book
	if err := entity.DB().Raw("SELECT * FROM books").Scan(&books).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}
// DELETE /books/:id
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM books WHERE id = ?", id); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}
// PATCH /books
func UpdateBook(c *gin.Context) {
	var book entity.Book
	if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if tx := entity.DB().Where("id = ?", book.ID).First(&book); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
			return
	}
	if err := entity.DB().Save(&book).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}