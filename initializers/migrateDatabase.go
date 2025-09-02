package initializers

import (
	"log"

	"example.com/m/v2/models"
)

func MigrateDatabase() {

	if err := DB.AutoMigrate(&models.Post{}); err != nil {
		log.Fatalf("Failed to migrate Post model: %v", err) //logs the error and stops the app.
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate User model: %v", err)
	}

	log.Println("Database migrated")
}
