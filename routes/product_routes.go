package routes

import (
	"freshers-bootcamp/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.POST("/product", controllers.AddProduct)
	r.PATCH("/product/:id", controllers.UpdateProduct)
	r.GET("/product/:id", controllers.GetProduct)
	r.GET("/products", controllers.GetAllProducts)
}
