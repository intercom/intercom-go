package intercom

// IntercomError is a known error from the Intercom API
type IntercomError interface {
	Error() string
	GetStatusCode() int
	GetCode() string
	GetMessage() string
}
