package controllers

import (
	"errors"
	"gin-training/database"
	"gin-training/helpers"
	"gin-training/models"
	"gin-training/serializers/requests"
	"gin-training/serializers/responses"
	"gin-training/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Register(ctx *gin.Context) {
	req := requests.RegisterRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]responses.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = responses.ErrorMsg{Field: fe.Field(), Message: fe.Error()}
			}
			response := helpers.BuildErrorResponse("Failed to register user", out)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		return
	}

	// Hash password
	user := req.BindRequest()
	if errPwd := user.HashPassword(user.Password); errPwd != nil {
		response := helpers.BuildErrorResponse("Failed to register user", errPwd.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Check duplicate email address
	if services.IsDuplicateEmail(user.Email) {
		response := helpers.BuildErrorResponse("Failed to register user", responses.ErrorMsg{Field: "email", Message: "Duplicate email address"})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create user
	if err := database.Instance.Model(&user).Create(&user).Error; err != nil {
		response := helpers.BuildErrorResponse("Failed to register user", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(user.Email, user.ID)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to register user", err.Error)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildDetailResponse(user.BindAuthResponse(token))
	ctx.JSON(201, &response)
}

func Login(ctx *gin.Context) {
	req := requests.LoginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]responses.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = responses.ErrorMsg{Field: fe.Field(), Message: fe.Error()}
			}
			response := helpers.BuildErrorResponse("Failed to login", out)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		return
	}

	// Find user by email
	var user models.User
	if err := database.Instance.Where("email = ?", req.Email).First(&user).Error; err != nil {
		response := helpers.BuildErrorResponse("Failed to login", "Invalid credentials")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// Check password
	credentialError := user.CheckPassword(req.Password)
	if credentialError != nil {
		response := helpers.BuildErrorResponse("Failed to login", "Invalid credentials")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// Generate JWT token
	token, errToken := helpers.GenerateJWT(user.Email, user.ID)
	if errToken != nil {
		response := helpers.BuildErrorResponse("Failed to login", errToken.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildDetailResponse(user.BindAuthResponse(token))
	ctx.JSON(200, &response)

}
