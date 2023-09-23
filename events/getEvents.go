package events

import (
	"net/http"

	"github.com/SayatAbdikul/innovatex_rest_api/server"
	"github.com/gin-gonic/gin"
)

type Event struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Date  string `json:"date"`
	Time  string `json:"time"`
	Image string `json:"image"`
}

func GetEvents(c *gin.Context) {
	server.Connect()
	defer server.DB.Close()
	rows, err := server.DB.Query("SELECT * FROM events")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		return
	}
	defer rows.Close()
	var events []Event

	// Iterate through the rows and scan data into the `discounts` slice
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Date, &event.Time, &event.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		events = append(events, event)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating rows"})
		return
	}

	// Send the discounts slice as a JSON response
	c.JSON(http.StatusOK, events)

}
