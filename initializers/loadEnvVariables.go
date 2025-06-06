package initializers

import (
	"fmt"
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
		log.Fatal("Error in filepath.Abs")
	}
	environmentPath := filepath.Join(dir, ".env")

	// Solve the .env file not loading error
	err = godotenv.Load(environmentPath)
	/*if err != nil {
		log.Fatal("Error loading the .env file")
	}
	*/
	fmt.Println(".env loaded successfully - loadEnvVariables.go")

	/*if err != nil {
		log.Fatal("Error loading the .env file")
	}*/
}
