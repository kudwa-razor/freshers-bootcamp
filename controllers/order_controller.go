package controllers

import (
	"freshers-bootcamp/databaseConnect"
	"freshers-bootcamp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	databaseConnect.DB.Create(&order)
	c.JSON(http.StatusOK, order)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := databaseConnect.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}
