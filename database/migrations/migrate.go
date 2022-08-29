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
	database.Instance.AutoMigrate(&models.User{})
}
