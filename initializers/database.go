//Using postgres for the database

package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// ConnectToDb Connects to the database - postgres
func ConnectToDb() {
	dsn :=
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(DB)
	if err != nil {
		log.Fatal("Error opening the database.")
	}

}
