package main

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {

	err := initializers.DB.AutoMigrate(&models.Post{})

	if err != nil {
		log.Fatal("Error in initializers.DB.AutoMigrate ")
	}

}
