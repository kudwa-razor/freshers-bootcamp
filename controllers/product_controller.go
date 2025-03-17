package controllers

import (
	"freshers-bootcamp/databaseConnect"
	"freshers-bootcamp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	databaseConnect.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := databaseConnect.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.BindJSON(&product)
	databaseConnect.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := databaseConnect.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	databaseConnect.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}
