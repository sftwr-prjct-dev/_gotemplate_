package utils

import (
	"context"

	"gitlab.com/coinprofile/services/gotemplate/src/config"
)

// RequestData ...
type RequestData interface {
	Validate() (isValid bool, errors []error)
	Controller(ctx context.Context, cfg *config.Config) (status int, msg string, data interface{}, err error)
	New() RequestData
	GetParamsMap() QueryMap
}

// Response general response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseErrors error structure
type ResponseErrors struct {
	Errors []string `json:"errors,omitempty"`
}

// CreateErrorResponse return string of errors from errors array
func CreateErrorResponse(errs ...error) (errStr []string) {

	for _, err := range errs {
		errStr = append(errStr, err.Error())
	}

	return
}

// QueryMap holds converted request query as map
type QueryMap map[string]interface{}
