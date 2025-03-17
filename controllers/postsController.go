package controllers

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

// PostsCreate - Creates the POSTS
func PostsCreate(c *gin.Context) {

	// Retrieve data off request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
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

// PostsIndex - PostsIndex
func PostsIndex(c *gin.Context) {

	//GET//Retrieve the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond accordingly with them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

// PostsShow - Retrieve Posts of a given number using the id
func PostsShow(c *gin.Context) {

	//Get id off url
	id := c.Param("id")
	//Retrieve the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with the found request
	c.JSON(200, gin.H{
		"post": post,
	})
}

// PostsUpdate - PostsUpdate
func PostsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")
	// Get the data off the req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post we're updating using the id
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

// PostsDelete - Deletes what has been POST/PUT
func PostsDelete(c *gin.Context) {

}
