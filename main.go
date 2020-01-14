package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ping Successful",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run())
}
