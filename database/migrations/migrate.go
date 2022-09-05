package main

/*
Purpose: Auto migrate models to database
Command: go run database/migrations/migrate.go
*/

import (
	"gin-training/database"
	"gin-training/models"
)

func init() {
	// Load environment variables
	database.LoadEnvVariables()
	// Open database connection
	database.SetupDatabaseConnection()
}

func main() {
	// Close database connection after application is closed
	defer database.CloseDatabaseConnection()

	// Auto migrate models to database
	database.Instance.AutoMigrate(
		&models.User{},
	)
}
