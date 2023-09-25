package missinglist

import (
	"net/http"

	"github.com/SayatAbdikul/innovatex_rest_api/server"
	"github.com/gin-gonic/gin"
)

type Missing struct {
	ID         int    `json:"id"`
	Category   string `json:"category"`
	Name       string `json:"name"`
	Additional string `json:"additional"`
	Image      string `json:"image"`
}

func GetMissingList(c *gin.Context) {
	server.Connect()
	defer server.DB.Close()
	rows, err := server.DB.Query("SELECT * FROM missingList")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		return
	}
	defer rows.Close()
	var list []Missing

	// Iterate through the rows and scan data into the `discounts` slice
	for rows.Next() {
		var missing Missing
		err := rows.Scan(&missing.ID, &missing.Category, &missing.Name, &missing.Additional, &missing.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		list = append(list, missing)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating rows"})
		return
	}

	// Send the discounts slice as a JSON response
	c.JSON(http.StatusOK, list)

}
