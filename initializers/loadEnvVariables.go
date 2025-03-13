package initializers

import (
	"log"
	"os"
)

func LoadEnvVariables() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
