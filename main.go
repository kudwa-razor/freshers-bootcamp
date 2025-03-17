package main

import (
	databaseConnect "freshers-bootcamp/databaseConnect"
	"freshers-bootcamp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	databaseConnect.ConnectDB()
	r := gin.Default()

	routes.ProductRoutes(r)
	routes.OrderRoutes(r)

	r.Run(":8080")
}
