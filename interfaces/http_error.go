package interfaces

import (
	"fmt"
	"net/http"
)

type HTTPErrorList struct {
	Type   string      `json:"type"`
	Errors []HTTPError `json:"errors"`
}

type HTTPError struct {
	StatusCode int
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func NewUnknownHTTPError(statusCode int) HTTPError {
	message := http.StatusText(statusCode)
	if message == "" {
		message = "Unknown Error"
	}
	return HTTPError{Code: "Unknown", Message: message, StatusCode: statusCode}
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%d: %s, %s", e.StatusCode, e.Code, e.Message)
}

func (e HTTPError) GetStatusCode() int {
	return e.StatusCode
}

func (e HTTPError) GetCode() string {
	return e.Code
}

func (e HTTPError) GetMessage() string {
	return e.Message
}
