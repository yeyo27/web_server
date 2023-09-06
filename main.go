package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "hello " + name,
	})
}

type Person struct {
	XMLName xml.Name `xml:"person"`
	FirstName string `xml:"firstName,attr"`
	LastName string `xml:"lastName,attr"`
}

func IndexHandlerXML(c *gin.Context) {
	name := c.Params.ByName("name")
	c.XML(200, Person{
		FirstName: name,
		LastName: "mesa",
	})
}


func main() {
	router := gin.Default()
	router.GET("/:name", IndexHandler)
	router.GET("/:name/xml", IndexHandlerXML)
	router.Run()
}