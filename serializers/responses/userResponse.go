package responses

type UserResponse struct {
	BaseResponse
	Name  string `json:"name"`
	Email string `json:"email"`
}
