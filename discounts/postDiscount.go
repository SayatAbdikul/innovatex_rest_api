package discounts

import (
	"net/http"

	"github.com/SayatAbdikul/innovatex_rest_api/server"
	"github.com/gin-gonic/gin"
)

func PostDiscount(c *gin.Context) {
	server.Connect()
	defer server.DB.Close()
	var product Discount
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err error
	_, err = server.DB.Exec("INSERT INTO discounts (title, oldprice, newprice) VALUES ($1, $2, $3)", product.Name, product.OldPrice, product.NewPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})

}
