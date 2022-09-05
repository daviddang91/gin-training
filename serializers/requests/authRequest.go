package requests

import "gin-training/models"

type RegisterRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
}

func (reg *RegisterRequest) BindRequest() models.User {
	user := models.User{
		Name:     reg.Name,
		Email:    reg.Email,
		Password: reg.Password,
	}
	return user
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
}
