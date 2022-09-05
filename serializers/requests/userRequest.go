package requests

import "gin-training/models"

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (reg *UpdateUserRequest) BindRequest(user *models.User) models.User {
	userObj := models.User{
		Name:     reg.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return userObj
}
