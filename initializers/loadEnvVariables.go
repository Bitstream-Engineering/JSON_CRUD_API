package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnvVariables loads the .env file
func LoadEnvVariables() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
