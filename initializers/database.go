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
	//dsn := os.Getenv("DB_URL")

	fmt.Println("Above dsn := ")
	dsn := "host=localhost user=postgres dbname=json_crud_api port=5432 sslmode=disable"
	fmt.Println("Middle of dsn & DB,err")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error opening the database.")
	}
	fmt.Println("Printing DB", DB)

}

// sudo -u postgres psql
// postgres dir /etc/postgresql/14/main
