//Using postgres for the database

package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

// ConnectToDb Connects to the database - postgres
func ConnectToDb() {
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(DB)
	if err != nil {
		log.Fatal("Error opening the database.")
	}

}
