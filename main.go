package main

import (
	"time"
	"github.com/gin-gonic/gin"
)


func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "hello " + name,
	})
}


type Recipe struct {
	Name string `json:"name"`
	Tags []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}


func main() {
	router := gin.Default()
	router.GET("/:name", IndexHandler)
	router.Run()
}