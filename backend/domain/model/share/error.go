package share

import "fmt"

type BusinessError interface {
	GetCodeError() string
	Error() string
	String() string
}

type ErrorExternalProvider struct {
	Code          string
	Message       string
	OriginalError string
}

func ProduceErrorExternalProvider(nameProvider, errorMsg string) BusinessError {
	return &ErrorExternalProvider{
		Code:          "EXTERNAL-001",
		Message:       fmt.Sprintf("has not been connect with the external service (%s)", nameProvider),
		OriginalError: errorMsg,
	}
}

func (error *ErrorExternalProvider) GetCodeError() string {
	return error.Code
}

func (error *ErrorExternalProvider) Error() string {
	return error.Message
}

func (error *ErrorExternalProvider) String() string {
	return fmt.Sprintf("ErrorExternalProvider %s - %s [%s]", error.Code, error.Message, error.OriginalError)
}
