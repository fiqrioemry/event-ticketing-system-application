package dto

// Response represents standard API response
type Response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Operation successful"`
	Data    any    `json:"data,omitempty"`
}

// ResponseWithPagination represents paginated API response
type ResponseWithPagination struct {
	Success    bool       `json:"success" example:"true"`
	Message    string     `json:"message" example:"Operation successful"`
	Data       any        `json:"data,omitempty"`
	Pagination Pagination `json:"pagination"`
}

// Pagination represents pagination information
type Pagination struct {
	Page       int `json:"page" example:"1"`
	Limit      int `json:"limit" example:"10"`
	Total      int `json:"total" example:"100"`
	TotalPages int `json:"total_pages" example:"10"`
}

// ErrorResponse represents error response
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Error occurred"`
	Error   string `json:"error,omitempty"`
}
