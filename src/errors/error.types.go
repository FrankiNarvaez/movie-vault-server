package errors

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Type       string `json:"type"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e CustomError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}

func NewError(typeError string, message string, code int) CustomError {
	return CustomError{
		Type:       typeError,
		Message:    message,
		StatusCode: code,
	}
}

func NewNotFoundError(message string) CustomError {
	return CustomError{
		StatusCode: http.StatusNotFound,
		Message:    message,
		Type:       "NOT_FOUND",
	}
}

func NewBadRequestError(message string) CustomError {
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Type:       "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) CustomError {
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Type:       "INTERNAL_SERVER_ERROR",
	}
}

func NewExternalAPIError(message string) CustomError {
	return CustomError{
		StatusCode: http.StatusBadGateway,
		Message:    message,
		Type:       "EXTERNAL_API_ERROR",
	}
}
