package migrations

import (
	"gin-training/database"
	"gin-training/models"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectDB()
}

func main() {
	database.DB.AutoMigrate(&models.User{})
}
