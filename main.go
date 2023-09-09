package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)


type Recipe struct {
	ID			 string `json:"id"`
	Name 		 string `json:"name"`
	Tags 		 []string `json:"tags"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipesDB []Recipe

func init() {
	recipesDB = make([]Recipe, 0)
}


func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hosting recipes app",
	})
}


func GetRecipesHandler(c *gin.Context) {

}


func PostRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipesDB = append(recipesDB, recipe)
	c.JSON(http.StatusOK, recipe)
}


func PutRecipesHandler(c *gin.Context) {

}

func DeleteRecipesHandler(c *gin.Context) {

}


func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.GET("/recipes", GetRecipesHandler)
	router.POST("/recipes", PostRecipeHandler)
	router.PUT("/recipes/:id", PutRecipesHandler)
	router.DELETE("/recipes/:id", DeleteRecipesHandler)
	router.Run()
}