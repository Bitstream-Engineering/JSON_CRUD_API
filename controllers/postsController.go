package controllers

import "C"
import "github.com/gin-gonic/gin"

func (c *gin.Context){
	c.JSON(200,  map[string]any{
		"message": "pong",
	})
}