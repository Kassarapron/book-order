package handler

import (
	"book-order-be/entity"
	"book-order-be/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBookOrder(c *gin.Context) {
	var req struct {
		Quantity  uint `json:"Quantity"`
		AdminId   uint `json:"AdminId"`
		CompanyId uint `json:"CompanyId"`
		BookId    uint `json:"BookId"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	bookOrder := entity.BookOrder{
		Quantity:  req.Quantity,
		AdminID:   &req.AdminId,
		CompanyID: &req.CompanyId,
		BookID:    &req.BookId,
	}

	if err := repository.InsertBookOrder(&bookOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, bookOrder)
}

func GetAllBookOrders(c *gin.Context) {
	res, err := repository.ListBookOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// func GetBookOrderById(c *gin.Context) {
// 	id := c.Param("id")

// 	res, err := repository.FindAdminById(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }
