package errors

import "net/http"

// RestErr struct
type RestErr struct {
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	Error      string `json:"error,omitempty"`
}

// NewBadRequestError returns a bad_request code error with string message error
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

// NewNotFoundError returns a not_found code error with string message error
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}

// NewInternalServerError returns a not_found code error with string message error
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
	}
}
