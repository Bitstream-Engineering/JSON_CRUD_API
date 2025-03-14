package controllers

import "C"
import "github.com/gin-gonic/gin"

func PostsCreate(c *gin.Context) {
	c.JSON(200, map[string]any{
		"message": "pong",
	})
}
