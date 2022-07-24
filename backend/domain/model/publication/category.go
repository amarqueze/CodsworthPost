package publication

import (
	"fmt"
	"strings"

	"outergeekhub.com/codsworthpost/domain/model/share"
)

type Category struct {
	Name string
}

func NewCategory(name string) *Category {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		panic(ProduceErrorPostState("Category's name is required"))
	}

	if len(name) < 3 {
		panic(ProduceErrorPostState("Category's name must be more 2 characters"))
	}

	return &Category{Name: name}
}

type InvalidCategoryState struct {
	Code          string
	Message       string
	OriginalError string
}

func ProduceErrorCategoryState(errorMsg string) share.BusinessError {
	return &InvalidCategoryState{
		Code:          "C-100",
		Message:       errorMsg,
		OriginalError: "Category Invalid State",
	}
}

func (error *InvalidCategoryState) GetCodeError() string {
	return error.Code
}

func (error *InvalidCategoryState) Error() string {
	return error.Message
}

func (error *InvalidCategoryState) String() string {
	return fmt.Sprintf("InvalidCategoryState %s - %s [%s]", error.Code, error.Message, error.OriginalError)
}

type CreationCategoryFailed struct {
	Code    string
	Message string
}

func ProduceErrorCreationCategory(errorMsg string) share.BusinessError {
	return &CreationCategoryFailed{
		Code:    "C-101",
		Message: "Category cannot created",
	}
}

func (error *CreationCategoryFailed) GetCodeError() string {
	return error.Code
}

func (error *CreationCategoryFailed) Error() string {
	return error.Message
}

func (error *CreationCategoryFailed) String() string {
	return fmt.Sprintf("CreationCategoryFailed %s - %s", error.Code, error.Message)
}

type CategoryNotFound struct {
	Code          string
	Message       string
	OriginalError string
}

func ProduceErrorCategoryNotFound(postId string) share.BusinessError {
	return &CategoryNotFound{
		Code:          "C-104",
		Message:       fmt.Sprintf("Post with Id %s not found", postId),
		OriginalError: "Category service not found Category",
	}
}

func (error *CategoryNotFound) GetCodeError() string {
	return error.Code
}

func (error *CategoryNotFound) Error() string {
	return error.Message
}

func (error *CategoryNotFound) String() string {
	return fmt.Sprintf("CategoryNotFound %s - %s [%s]", error.Code, error.Message, error.OriginalError)
}
