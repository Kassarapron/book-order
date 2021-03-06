package controller
import (
        "github.com/kassarapron/sa-64/entity"
        "github.com/gin-gonic/gin"
        "net/http"
)
// POST /companys
func CreateCompany(c *gin.Context) {
	var company entity.Company
	if err := c.ShouldBindJSON(&company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if err := entity.DB().Create(&company).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": company})
}
// GET /company/:id
func GetCompany(c *gin.Context) {
	var company entity.Company
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM companys WHERE id = ?", id).Scan(&company).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": company})
}
// GET /companys
func ListCompanys(c *gin.Context) {
	var companys []entity.Company
	if err := entity.DB().Raw("SELECT * FROM companys").Scan(&companys).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": companys})
}
// DELETE /companys/:id
func DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM companys WHERE id = ?", id); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "company not found"})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}
// PATCH /companys
func UpdateCompany(c *gin.Context) {
	var company entity.Company
	if err := c.ShouldBindJSON(&company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if tx := entity.DB().Where("id = ?", company.ID).First(&company); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "company not found"})
			return
	}
	if err := entity.DB().Save(&company).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": company})
}