package main

import (
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}


func main() {
	router := gin.Default()
	router.GET("/", indexHandler)
	router.Run()
}