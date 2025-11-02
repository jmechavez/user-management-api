package apperrors

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:"errorCode,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewInvalidUserIDError() *AppError {
	return &AppError{
		Message: "Invalid user ID",
		Code:    http.StatusBadRequest,
	}
}
