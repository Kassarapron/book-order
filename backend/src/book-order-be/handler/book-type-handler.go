package handler

import (
	"book-order-be/entity"
	"book-order-be/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBookType(c *gin.Context) {
	var req struct {
		TypeName string `json:"TypeName"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	bookType := entity.BookType{
		TypeName: req.TypeName,
	}
	if err := repository.InsertBookType(&bookType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, bookType)
}

func GetAllBookTypes(c *gin.Context) {
	res, err := repository.ListBookTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// func GetBookTypeById(c *gin.Context) {
// 	id := c.Param("id")

// 	res, err := repository.FindBookTypeById(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }
