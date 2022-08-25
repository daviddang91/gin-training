package main

import (
	"gin-training/database"
	"gin-training/routes"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectDB()
}

func main() {
	routes.Run()
}
