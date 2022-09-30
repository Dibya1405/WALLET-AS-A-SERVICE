package dto

type Response struct {
	StatusCode int `json:"status_code"`
	Id         int `json:"id"`
}

type SuccessfulResponse struct {
	Response
	SuccessMessage string `json:"success_message"`
}

type ErrorResponse struct {
	Response
	ErrorMessage string `json:"error_message"`
}
