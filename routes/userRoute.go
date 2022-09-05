package routes

import (
	"gin-training/controllers"
	"gin-training/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.Use(middlewares.Authenticate())

	users.GET("/", controllers.ListUser)
	users.GET("/:id", controllers.DetailUser)
	users.POST("/", controllers.CreateUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)
}
