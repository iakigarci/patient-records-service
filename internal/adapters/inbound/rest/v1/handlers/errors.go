package handlers

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid credentials"`
}
