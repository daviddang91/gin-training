package responses

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
