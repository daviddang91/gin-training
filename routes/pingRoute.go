package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
}
