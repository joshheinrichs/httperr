/*
Package httperr defines a simple HTTP error which has a message and status
code.
*/
package httperr

// Error represents a simple HTTP error which has a message and status code.
type Error interface {
	error
	// Code returns the error's HTTP status code.
	Code() int
}

type httpError struct {
	message string
	code    int
}

// New returns a new Error with the given message and code.
func New(message string, code int) Error {
	return &httpError{message, code}
}

// Error returns the error's message.
func (e *httpError) Error() string {
	return e.message
}

// Code returns the error's HTTP status code.
func (e *httpError) Code() int {
	return e.code
}
