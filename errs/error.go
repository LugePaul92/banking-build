package errs

import "net/http"

type AppError struct{
	Code int		`json:",omitempty"`
	Message string	`json:"message"`
}

func (a AppError) AsMessage() *AppError{
	return &AppError{
		Message: a.Message,
	}
}
func NewNotFoundError(message string) *AppError{
	return &AppError{
		Code: http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message string) *AppError{
	return &AppError{
		Code: http.StatusInternalServerError,
		Message: message,
	}
}

func NewValidationError(message string) *AppError{
	return &AppError{
		Code: http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}