package main

import (
	"fmt"

	"github.com/SayatAbdikul/innovatex_api/server"
	"github.com/gin-gonic/gin"
)

func dFunc(c *gin.Context) {
	fmt.Fprintf(c.Writer, "something")
}
func main() {
	server.Connect()
	defer server.DB.Close()
	router := gin.Default()
	router.GET("/", dFunc)
	router.Run(":8080")

}
