package main

import (
	"gin-training/routes"
)

func main() {
	// Our server will live in the routes package
	routes.Run()
}
