package responses

type AuthResponse struct {
	BaseResponse
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
