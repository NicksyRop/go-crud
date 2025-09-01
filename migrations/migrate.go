package main

import (
	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

// create a table automatically in our db using migration
// run go run fileName i.e go run migrations/migrate.go
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
