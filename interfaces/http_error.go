package interfaces

import "fmt"

type HTTPErrorList struct {
	Type   string      `json:"type"`
	Errors []HTTPError `json:"errors"`
}

type HTTPError struct {
	StatusCode int
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func NewUnknownHTTPError() HTTPError {
	return HTTPError{Code: "Unknown", Message: "Unknown Error"}
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
