package discounts

import (
	"net/http"

	"github.com/SayatAbdikul/innovatex_rest_api/server"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func PatchDiscount(c *gin.Context) {
	// Parse the ID parameter from the URL

	// Parse the JSON data from the request body into a Product struct
	var updatedProduct Discount
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	server.Connect()
	defer server.DB.Close()
	var err error
	// Execute an UPDATE SQL statement to update the data
	_, err = server.DB.Exec("UPDATE products SET (title = $1, oldprice = $2, newprice = $3, image = $4) WHERE id = $5",
		updatedProduct.Name, updatedProduct.OldPrice, updatedProduct.NewPrice, updatedProduct.Image, updatedProduct.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
