package database

import "gin-training/models"

func init() {
	LoadEnvVariables()
	ConnectDB()
}

func main() {
	DB.AutoMigrate(&models.User{})
}
