package discounts

import (
	"net/http"

	"github.com/SayatAbdikul/innovatex_rest_api/server"
	"github.com/gin-gonic/gin"
)

type Discount struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	OldPrice int    `json:"oldprice"`
	NewPrice int    `json:"newprice"`
	Image    string `json:"image"`
}

func GetDiscounts(c *gin.Context) {
	server.Connect()
	defer server.DB.Close()
	rows, err := server.DB.Query("SELECT * FROM discounts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		return
	}
	defer rows.Close()
	var discounts []Discount

	// Iterate through the rows and scan data into the `discounts` slice
	for rows.Next() {
		var discount Discount
		x := &discount.Image
		err := rows.Scan(&discount.Name, &discount.OldPrice, &discount.NewPrice, &discount.ID, x)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		discounts = append(discounts, discount)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating rows"})
		return
	}

	// Send the discounts slice as a JSON response
	c.JSON(http.StatusOK, discounts)

}
