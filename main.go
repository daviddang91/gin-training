package main

import (
	"gin-training/database"
	"gin-training/routes"
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
	// Start application
	routes.Run()
}
