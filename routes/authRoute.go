package routes

import (
	"gin-training/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(auth *gin.RouterGroup) {
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
}
