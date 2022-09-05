package requests

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

// func (reg *UpdateUserRequest) BindRequest(user *models.User) models.User {
// 	userObj := models.User{
// 		Name:     reg.Name,
// 		Email:    user.Email,
// 		Password: reg.Password,
// 	}
// 	return userObj
// }
