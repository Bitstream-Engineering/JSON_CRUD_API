package main

import (
	"JSON_CRUD_API/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()

}

func main() {

	// GIN boilerplate code
	//gin.SetMode(gin.ReleaseMode) //Set for production use
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
