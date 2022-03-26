package spot

import (
	"fmt"
	"net/http"
	"strings"
)

type ClientError struct {
	// https status code
	StatusCode int64
	// error code returned from server
	ErrorCode int64
	// error message returned from server
	ErrorMessage string
	// the whole response header returned from server
	Header http.Header
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("client error::status code: %d, error code: %d, message: %s", e.StatusCode, e.ErrorCode, e.ErrorMessage)
}

type ServerError struct {
	StatusCode int64
	Message    string
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("server error::status code: %d, message: %s", e.StatusCode, e.Message)
}

type ParameterRequiredError struct {
	// Params ([]string): key of parameter
	Params []string
}

func (e *ParameterRequiredError) Error() string {
	return fmt.Sprintf("%s is mandatory, but received empty.", strings.Join(e.Params, ", "))
}

type ParameterValueError struct {
	// Params ([]string): key of parameter
	Params []string
}

func (e *ParameterValueError) Error() string {
	return fmt.Sprintf("the enum value %s is invalid.", strings.Join(e.Params, ", "))
}

type ParameterArgumentError struct {
	ErrorMessage string
}

func (e *ParameterArgumentError) Error() string {
	return e.ErrorMessage
}
