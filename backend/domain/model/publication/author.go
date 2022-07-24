package publication

import (
	"fmt"
	"strings"

	"outergeekhub.com/codsworthpost/domain/helper"
	"outergeekhub.com/codsworthpost/domain/model/share"
)

type Author struct {
	Id          string
	Fullname    string
	Description string
}

func NewAuthor(fullname, description string) *Author {
	fullname = strings.TrimSpace(fullname)
	if len(fullname) == 0 {
		panic(ProduceErrorPostState("Author's name is required"))
	}

	id := helper.Util{}.GenerateId()
	description = strings.TrimSpace(description)
	return &Author{Id: id, Fullname: fullname, Description: description}
}

type InvalidAuthorState struct {
	Code    string
	Message string
}

func ProduceErrorAuthorState(errorMsg string) share.BusinessError {
	return &InvalidAuthorState{
		Code:    "A-100",
		Message: errorMsg,
	}
}

func (error *InvalidAuthorState) GetCodeError() string {
	return error.Code
}

func (error *InvalidAuthorState) Error() string {
	return error.Message
}

func (error *InvalidAuthorState) String() string {
	return fmt.Sprintf("InvalidAuthorState %s - %s", error.Code, error.Message)
}

type CreationAuthorFailed struct {
	Code          string
	Message       string
	OriginalError string
}

func ProduceErrorCreationAuthor(errorMsg string) share.BusinessError {
	return &CreationAuthorFailed{
		Code:          "A-101",
		Message:       "Author cannot created",
		OriginalError: errorMsg,
	}
}

func (error *CreationAuthorFailed) GetCodeError() string {
	return error.Code
}

func (error *CreationAuthorFailed) Error() string {
	return error.Message
}

func (error *CreationAuthorFailed) String() string {
	return fmt.Sprintf("CreationPostFailed %s - %s [%s]", error.Code, error.Message, error.OriginalError)
}

type AuthorNotFound struct {
	Code          string
	Message       string
	OriginalError string
}

func ProduceErrorAuthorNotFound(postId string) share.BusinessError {
	return &AuthorNotFound{
		Code:          "A-104",
		Message:       fmt.Sprintf("Post with Id %s not found", postId),
		OriginalError: "Category service not found Category",
	}
}

func (error *AuthorNotFound) GetCodeError() string {
	return error.Code
}

func (error *AuthorNotFound) Error() string {
	return error.Message
}

func (error *AuthorNotFound) String() string {
	return fmt.Sprintf("CategoryNotFound %s - %s [%s]", error.Code, error.Message, error.OriginalError)
}
