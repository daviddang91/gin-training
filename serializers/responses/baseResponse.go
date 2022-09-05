package responses

import (
	"time"

	"gorm.io/gorm"
)

type BaseResponse struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type DetailResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ListResponse struct {
	Status     bool        `json:"status"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

type ErrorResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
