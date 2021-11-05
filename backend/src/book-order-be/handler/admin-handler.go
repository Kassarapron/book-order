package handler

import (
	"book-order-be/entity"
	"book-order-be/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAdmin(c *gin.Context) {
	var req struct {
		AdminName string `json:"AdminName"`
		Email     string `json:"Email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	admin := entity.Admin{
		AdminName: req.AdminName,
		Email:     req.Email,
	}
	if err := repository.InsertAdmin(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, admin)
}

func GetAllAdmins(c *gin.Context) {
	res, err := repository.ListAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetAdminById(c *gin.Context) {
	id := c.Param("id")

	res, err := repository.FindAdminById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
