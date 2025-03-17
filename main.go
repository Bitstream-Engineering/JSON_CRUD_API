package main

import (
	"JSON_CRUD_API/controllers"
	"JSON_CRUD_API/initializers"
	"github.com/gin-gonic/gin"
)

// init Initializes multiple functions
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {

	// GIN boilerplate code
	//gin.SetMode(gin.ReleaseMode) //Set for production use
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	err := r.Run()
	if err != nil {
		return
	}
}
