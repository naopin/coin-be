package util

import "fmt"

type ErrorCode string

const (
	ErrorCodeRequired            ErrorCode = "1001"
	ErrorCodeUnexpectedSituation ErrorCode = "9999"
)

type AppError struct {
	Code    ErrorCode
	Message string
}

func (e AppError) Error() string {
	return fmt.Sprintf("[%s]%s", e.Code, e.Message)
}
