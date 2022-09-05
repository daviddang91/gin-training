package middlewares

import (
	"gin-training/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			response := helpers.BuildErrorResponse("Unauthorized", "Request does not contain an access token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		parts := strings.SplitN(tokenString, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response := helpers.BuildErrorResponse("Unauthorized", "Request header auth Incorrect format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		err := helpers.ValidateToken(parts[1])
		if err != nil {
			response := helpers.BuildErrorResponse("Unauthorized", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Next()
	}
}
