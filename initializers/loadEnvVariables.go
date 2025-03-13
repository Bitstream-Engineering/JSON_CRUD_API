package initializers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

// LoadEnvVariables loads the .env file
func LoadEnvVariables() {
	//err := godotenv.Load(".env")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")

	godotenv.Load(environmentPath)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
