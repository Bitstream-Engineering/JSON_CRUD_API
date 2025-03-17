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
	r.POST("/posts", controllers.PostsCreate)    // POST is a safe method because it doesn't modify the state of the server
	r.PUT("/posts/:id", controllers.PostsUpdate) // PUT updates or replaces existing data on server
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("posts/:id", controllers.PostsDelete)

	err := r.Run()
	if err != nil {
		return
	}
}
