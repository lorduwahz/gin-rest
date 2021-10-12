package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lorduwahz/gin-rest/products"
)

// This is the handler that displays all our products.
func GetProducts(c *gin.Context) {
	prodList := products.GetProducts()
	c.IndentedJSON(http.StatusOK, prodList)
}

// This is the handler that displays specific product when it id is provided.
func GetProductsByID(c *gin.Context) {
	// a GET request to /user/john
	id, _ := strconv.Atoi(c.Param("id"))
	product, _, err := products.GetProdID(id)
	if err != nil {
		fmt.Println("Id is not in our database.")
	}
	c.IndentedJSON(http.StatusOK, product)
}

// This handler adds a new product to our list then displays the new list.
func AddProduct(c *gin.Context) {

	var newProduct *products.Product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products.AddProd(newProduct)

	c.IndentedJSON(http.StatusCreated, newProduct)

	/* 	A function to print our post result to file.

		func prodTofile(prodList []*Product) {
	 		file, err := os.Create("product.txt")

	 		if err != nil {
	 			fmt.Println("Can not create product file")
	 		}
	 		fmt.Fprintln(file, prodList)
	 	}
	*/
}

// This handler deletes a product when its id is provided.
func DelProd(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	newProdList := products.DeleteProd(id)
	c.IndentedJSON(http.StatusOK, newProdList)
}
