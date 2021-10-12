package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lorduwahz/gin-rest/handlers"
)

func main() {
	// Initialising our router
	r := gin.Default()

	// Routes with paths and their corresponding handler functions
	r.GET("/", handlers.GetProducts)
	r.GET("/product/:id", handlers.GetProductsByID)
	r.POST("/addProduct", handlers.AddProduct)
	r.DELETE("/product/:id", handlers.DelProd)

	// Attaches our routers to http.Server to listen at port: 9090.
	r.Run("localhost:9090")
}
