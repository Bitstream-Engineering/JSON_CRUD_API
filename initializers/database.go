//Using postgres for the database

package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// ConnectToDb Connects to the database - postgres
func ConnectToDb() {
	//dsn := os.Getenv("DB_URL")

	dsn := "host=localhost user=postgres dbname=json_crud_api port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//fmt.Println(DB)
	if err != nil {
		log.Fatal("Error opening the database.")
	}

}

// sudo -u postgres psql
// postgres dir /etc/postgresql/14/main
