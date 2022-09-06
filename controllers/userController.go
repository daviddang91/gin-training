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

func ListUser(ctx *gin.Context) {
	var model models.User
	var users []responses.UserResponse

	p := helpers.GeneratePagination(ctx, model)
	pagination, err := helpers.Paginator(&p, &users)

	if err != nil {
		response := helpers.BuildErrorResponse("Failed to get list of users", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildListResponse(&users, &pagination)
	ctx.JSON(200, &response)
}

func DetailUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	var user models.User

	if err := database.Instance.Model(&models.User{}).First(&user, userId).Error; err != nil {
		response := helpers.BuildErrorResponse("Failed to get user detail", err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response := helpers.BuildDetailResponse(user.BindResponse())
	ctx.JSON(200, &response)
}

func CreateUser(ctx *gin.Context) {
	req := requests.RegisterRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]responses.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = responses.ErrorMsg{Field: fe.Field(), Message: fe.Error()}
			}
			response := helpers.BuildErrorResponse("Failed to process request", out)
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
		response := helpers.BuildErrorResponse("Failed to create user", responses.ErrorMsg{Field: "email", Message: "Duplicate email address"})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Create user
	if err := database.Instance.Model(&user).Create(&user).Error; err != nil {
		response := helpers.BuildErrorResponse("Failed to create user", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildDetailResponse(user.BindResponse())
	ctx.JSON(201, &response)
}

func UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	var user models.User

	// Find user
	if err := database.Instance.Model(&models.User{}).First(&user, userId).Error; err != nil {
		response := helpers.BuildErrorResponse("Failed to update user", "User not found")
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	// Bind data from json to user object
	req := requests.UpdateUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]responses.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = responses.ErrorMsg{Field: fe.Field(), Message: fe.Error()}
			}
			response := helpers.BuildErrorResponse("Failed to update user", out)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		return
	}

	// Update user
	user = req.BindRequest(&user)
	if err := database.Instance.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		response := helpers.BuildErrorResponse("Failed to update user", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildDetailResponse(user.BindResponse())
	ctx.JSON(200, &response)
}

func DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	var user models.User

	// Find user
	if err := database.Instance.Model(&models.User{}).First(&user, userId).Error; err != nil {
		response := helpers.BuildErrorResponse("Failed to delete user", "User not found")
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	// Delete user
	if err := database.Instance.Model(&models.User{}).Delete(&user).Error; err != nil {
		response := helpers.BuildErrorResponse("Failed to delete user", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	ctx.Status(204)
}
