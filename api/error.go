package api

import (
	"encoding/json"
)

type ApiError struct {
	Message string `json:"message"`
}

func NewApiError(message string) ApiError {
	return ApiError{
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