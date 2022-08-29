package routes

import (
	"gin-training/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(auth *gin.RouterGroup) {
	auth.POST("/signup", controllers.RegisterUser)
	auth.POST("/login", controllers.GenerateToken)
}
