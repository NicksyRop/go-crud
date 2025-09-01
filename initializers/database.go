package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//create global db variable

var DB *gorm.DB

func ConnectDatabase() {

	// gorm.Open returns DB (if successful ) and err  if something went wrong - otherwise nil
	var err error
	dsn := os.Getenv("DB_CONNECTION_STRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
