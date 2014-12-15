package intercom

import "fmt"

type HttpErrorList struct {
	Type   string      `json:"type"`
	Errors []HttpError `json:"errors"`
}

type HttpError struct {
	StatusCode int
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%d: %s, %s", e.StatusCode, e.Code, e.Message)
}

type UnknownHttpError struct {
	*HttpError
}

func NewUnknownHttpError(statusCode int) *UnknownHttpError {
	return &UnknownHttpError{HttpError: &HttpError{StatusCode: statusCode, Code: "Unknown", Message: "Unknown Error"}}
}
