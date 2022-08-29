package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	getRoutes()
	router.Run(serverAddress)
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	auth := router.Group("/auth")
	AuthRoutes(auth)

	v1 := router.Group("/v1")
	UserRoutes(v1)
	//PingRoutes(v1)

	// v2 := router.Group("/v2")
	// PingRoutes(v2)
}
