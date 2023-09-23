package main

import (
	"fmt"

	discounts "github.com/SayatAbdikul/innovatex_rest_api/Discounts"
	"github.com/SayatAbdikul/innovatex_rest_api/events"
	missinglist "github.com/SayatAbdikul/innovatex_rest_api/missingList"
	"github.com/SayatAbdikul/innovatex_rest_api/server"
	"github.com/gin-gonic/gin"
)

func dFunc(c *gin.Context) {
	fmt.Fprintf(c.Writer, "something")
}
func main() {
	server.Connect()
	defer server.DB.Close()
	router := gin.Default()
	router.Use(server.CORSMiddleware())
	router.GET("/", dFunc)
	router.GET("/discounts", discounts.GetDiscounts)
	router.GET("/events", events.GetEvents)
	router.GET("/missing", missinglist.GetMissingList)

	router.Run(":8080")

}
