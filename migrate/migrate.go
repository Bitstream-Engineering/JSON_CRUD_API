package migrate

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
