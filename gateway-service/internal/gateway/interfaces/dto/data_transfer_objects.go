package dto

type Payload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Response struct {
	StatusCode int              `json:"status_code"`
	Message    string           `json:"message"`
	Data       []map[string]any `json:"data"`
	Error      bool             `json:"error"`
}
