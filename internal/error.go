package internal

type ApiError struct {
	Message string `json:"message"`
}

func NewApiError(message string) *ApiError {
	return &ApiError{message}
}
