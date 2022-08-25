package routes

import (
	"gin-training/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", controllers.GetUsersController)
	users.GET("/:id", controllers.DetailUserController)
	users.POST("/", controllers.CreateUserController)
	users.PUT("/:id", controllers.UpdateUserController)
	users.DELETE("/:id", controllers.DeleteUserController)

	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})

	users.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}
