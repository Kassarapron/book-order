package controller
import (
        "github.com/kassarapron/sa-64/entity"
        "github.com/gin-gonic/gin"
        "net/http"
)
// POST /bookorders
func CreateBookOrder(c *gin.Context) {
	var bookorder entity.BookOrder
	if err := c.ShouldBindJSON(&bookorder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if err := entity.DB().Create(&bookorder).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookorder})
}
// GET /bookorder/:id
func GetBookOrder(c *gin.Context) {
	var bookorder entity.BookOrder
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM book_orders WHERE id = ?", id).Scan(&bookorder).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookorder})
}
// GET /bookorders
func ListBookOrders(c *gin.Context) {
	var bookorders []entity.BookOrder
	if err := entity.DB().Raw("SELECT * FROM book_orders").Scan(&bookorders).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookorders})
}
// DELETE /bookorders/:id
func DeleteBookOrder(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM book_orders WHERE id = ?", id); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookorder not found"})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}
// PATCH /bookorders
func UpdateBookOrder(c *gin.Context) {
	var bookorder entity.BookOrder
	if err := c.ShouldBindJSON(&bookorder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if tx := entity.DB().Where("id = ?", bookorder.ID).First(&bookorder); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookorder not found"})
			return
	}
	if err := entity.DB().Save(&bookorder).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookorder})
}