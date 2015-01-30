package client

type IntercomError interface {
	Error() string
	GetStatusCode() int
	GetCode() string
	GetMessage() string
}
