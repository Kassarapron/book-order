package handler

import (
	"book-order-be/entity"
	"book-order-be/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var req struct {
		BookName      string `json:"BookName"`
		BookNumber    string `json:"BookNumber"`
		BookAuthor    string `json:"BookAuthor"`
		BookPublicher string `json:"BookPublicher"`
		BookTypeID    uint   `json:"BookTypeID"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book := entity.Book{
		BookName:      req.BookName,
		BookNumber:    req.BookNumber,
		BookAuthor:    req.BookAuthor,
		BookPublicher: req.BookPublicher,
		BookTypeID:    &req.BookTypeID,
	}
	if err := repository.InsertBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func GetAllBooks(c *gin.Context) {
	res, err := repository.ListBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")

	res, err := repository.FindBookById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
