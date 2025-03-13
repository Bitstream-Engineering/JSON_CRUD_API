package main

import (
	"JSON_CRUD_API/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()

}

func main() {

	//gin.SetMode(gin.ReleaseMode) //Set for production use
	// GIN boilerplate code
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
