package api

import (
	"encoding/json"
)

type ApiError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewApiError(message string) ApiError {
	return ApiError{
		Success: false,
		Message: message,
	}
}

func (e *ApiError) Error() string {
	return string(e.Json())
}

func (e *ApiError) Json() []byte {
	val, err := json.Marshal(e);
	if err != nil {
		return []byte{}
	}
	return val
}
