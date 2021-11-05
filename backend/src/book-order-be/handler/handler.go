package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"book-order-be/entity"
	"book-order-be/repository"
)

func ListUsers(c *gin.Context) {
	var query struct {
		Age uint `form:"age"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := repository.FindUsers(query.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetUser(c *gin.Context) {
	var param struct {
		Name string `uri:"name"`
	}

	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := repository.FindByName(param.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      uint   `json:"age"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := entity.User{
		Username: req.Username,
		Password: req.Password,
		Name: req.Name,
		Age: req.Age,
	}
	if err := repository.InsertUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
