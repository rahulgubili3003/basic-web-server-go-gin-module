package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello. This a Json Response from the GO web server using Gin module",
		})
	})
	if err := r.Run(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
