# httperr [![](https://godoc.org/github.com/joshheinrichs/httperr?status.svg)](https://godoc.org/github.com/joshheinrichs/httperr)

httperr defines a simple HTTP error which has a message and status code. This 
can be useful for returning errors with relevant HTTP status codes, which the
standard errors package allow for.

Example:

```go
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
```
