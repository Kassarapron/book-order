package handler

import (
	"book-order-be/entity"
	"book-order-be/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context) {
	var req struct {
		CompanyName string `json:"CompanyName"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	company := entity.Company{
		CompanyName: req.CompanyName,
	}
	if err := repository.InsertCompany(&company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, company)
}

func GetAllCompanies(c *gin.Context) {
	res, err := repository.ListCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// func GetCompanyById(c *gin.Context) {
// 	id := c.Param("id")

// 	res, err := repository.FindCompanyById(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }
