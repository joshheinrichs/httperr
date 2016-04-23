/*
httperr defines a simple HTTP error which has a message and status code. This 
can be useful for returning errors with relevant HTTP status codes, which the
standard errors package allow for.

Example:
	func Foo() httperr.Error {
	  if err := Bar(); err != nil {
	    return httperr.New(err.Error(), http.StatusInternalServerError)
	  }
	  return nil
	}
	
	func FooHandler(w http.ResponseWriter, r *http.Request) {
	  if httpErr := Foo(); httpErr != nil {
	    http.Error(w, httpErr.Error(), httpErr.Code())
	  }
	}
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
