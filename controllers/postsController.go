package controllers

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"github.com/gin-gonic/gin"
	"log"
)

func PostsCreate(c *gin.Context) {

	post := models.Post{Title: "Title", Body: "Post Body"}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		log.Fatal("Error in initializers.DB.Create(&post) ")
	}

	c.JSON(200, map[string]any{
		"message": "pong",
	})
}
