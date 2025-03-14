package controllers

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Retrieve data off request body
	var body struct {
		Body  string
		Title string
	}

	// Create a post
	post := models.Post{
		Title: "Title",
		Body:  "Post Body",
	}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		//log.Fatal("Error in initializers.DB.Create(&post)")
		c.Status(400)
		return
	}

	fmt.Println("Rows affected:", result.RowsAffected)

	c.JSON(200, map[string]any{
		"post": post,
	})
}
