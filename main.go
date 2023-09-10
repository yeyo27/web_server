package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	fileBytes, err := os.ReadFile("recipes.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error occurred:", err)
		return
	}
	_ = json.Unmarshal(fileBytes, &recipesDB)
}


func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hosting recipes app",
	})
}


func GetRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipesDB)
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
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	index := -1
	for i := 0; i < len(recipesDB); i++ {
		if recipesDB[i].ID  == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})
		return
	}
	recipe.ID = recipesDB[index].ID
	recipe.PublishedAt = recipesDB[index].PublishedAt
	c.JSON(http.StatusOK, recipe)
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