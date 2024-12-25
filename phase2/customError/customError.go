package customError

import "fmt"

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("ERROR - Code: %d, Message: %s", e.Code, e.Message)
}

func NewMyError(code int, message string) *MyError {
	return &MyError{Code: code, Message: message}
}
