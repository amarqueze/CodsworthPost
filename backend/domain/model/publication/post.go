package publication

import (
	"fmt"
	"strings"
	"time"

	"outergeekhub.com/codsworthpost/domain/helper"
	"outergeekhub.com/codsworthpost/domain/model/share"
)

type Post struct {
	Id              string
	Title           string
	Summary         string
	Content         string
	PrimaryImageRef string
	Category        *Category
	Tags            []string
	DateCreated     string
	DateEdition     string
	Lang            string
	Author          *Author
	IsPrimary       bool
	Reference       []string
}

func NewPost(lang string) *NewPostBuilder {
	newPostBuilder := new(NewPostBuilder)
	newPostBuilder.post.Lang = lang
	newPostBuilder.post.IsPrimary = true
	return newPostBuilder
}

type NewPostBuilder struct {
	post Post
}

func (builder *NewPostBuilder) Title(title string) *NewPostBuilder {
	builder.post.Title = title
	return builder
}

func (builder *NewPostBuilder) Summary(summary string) *NewPostBuilder {
	builder.post.Summary = summary
	return builder
}

func (builder *NewPostBuilder) Content(content string) *NewPostBuilder {
	builder.post.Content = content
	return builder
}

func (builder *NewPostBuilder) PrimaryImage(primaryImageRef string) *NewPostBuilder {
	builder.post.PrimaryImageRef = primaryImageRef
	return builder
}

func (builder *NewPostBuilder) Category(category *Category) *NewPostBuilder {
	builder.post.Category = category
	return builder
}

func (builder *NewPostBuilder) AddTags(tags []string) *NewPostBuilder {
	builder.post.Tags = tags
	return builder
}

func (builder *NewPostBuilder) Writer(author *Author) *NewPostBuilder {
	builder.post.Author = author
	return builder
}

func (builder *NewPostBuilder) Primary(isPrimary bool) *NewPostBuilder {
	builder.post.IsPrimary = isPrimary
	return builder
}

func (builder *NewPostBuilder) Build() *Post {
	builder.post.Title = strings.TrimSpace(builder.post.Title)
	if len(builder.post.Title) == 0 {
		panic(ProduceErrorPostState("Title is required"))
	}
	if len(builder.post.Title) < 3 {
		panic(ProduceErrorPostState("Title must be more 2 characters"))
	}

	builder.post.Summary = strings.TrimSpace(builder.post.Summary)
	if len(builder.post.Summary) == 0 {
		panic(ProduceErrorPostState("Summary is required"))
	}

	builder.post.Content = strings.TrimSpace(builder.post.Content)
	if len(builder.post.Content) == 0 {
		panic(ProduceErrorPostState("Content is required"))
	}

	if builder.post.Category == nil {
		panic(ProduceErrorPostState("Category is required"))
	}

	if builder.post.Author == nil {
		panic(ProduceErrorPostState("Author is required"))
	}

	builder.post.Id = helper.Util{}.GenerateId()
	builder.post.DateCreated = time.Now().String()
	builder.post.DateEdition = time.Now().String()
	return &builder.post
}

func EditPost(id string) *EditPostBuilder {
	if len(id) == 0 {
		panic(ProduceErrorPostState("PostId is required"))
	}

	editPostBuilder := new(EditPostBuilder)
	editPostBuilder.post.Id = id
	return editPostBuilder
}

type EditPostBuilder struct {
	post Post
}

func (builder *EditPostBuilder) Content(content string) *EditPostBuilder {
	builder.post.Content = strings.TrimSpace(content)
	if len(builder.post.Content) == 0 {
		panic(ProduceErrorPostState("Content is required"))
	}

	return builder
}

func (builder *EditPostBuilder) AddTags(tags []string) *EditPostBuilder {
	builder.post.Tags = tags
	return builder
}

func (builder *EditPostBuilder) Build() *Post {
	builder.post.DateEdition = time.Now().String()
	return &builder.post
}

func AddReferencePost(postId, reference string) *Post {
	if len(postId) == 0 {
		panic(ProduceErrorPostState("PostId is required"))
	}

	if len(reference) == 0 {
		panic(ProduceErrorPostState("Post reference is required"))
	}

	post := new(Post)
	post.Id = postId
	post.Reference = append(post.Reference, reference)
	return post
}

type InvalidPostState struct {
	Code    string
	Message string
}

func ProduceErrorPostState(errorMsg string) share.BusinessError {
	return &InvalidPostState{
		Code:    "P-100",
		Message: errorMsg,
	}
}

func (error *InvalidPostState) GetCodeError() string {
	return error.Code
}

func (error *InvalidPostState) Error() string {
	return error.Message
}

func (error *InvalidPostState) String() string {
	return fmt.Sprintf("InvalidPostState %s - %s", error.Code, error.Message)
}

type CreationPostFailed struct {
	Code          string
	Message       string
	OriginalError string
}

func ProduceErrorCreationPost(errorMsg string) share.BusinessError {
	return &CreationPostFailed{
		Code:          "P-101",
		Message:       "Post cannot Publish",
		OriginalError: errorMsg,
	}
}

func (error *CreationPostFailed) GetCodeError() string {
	return error.Code
}

func (error *CreationPostFailed) Error() string {
	return error.Message
}

func (error *CreationPostFailed) String() string {
	return fmt.Sprintf("CreationPostFailed %s - %s [%s]", error.Code, error.Message, error.OriginalError)
}

type EditionPostFailed struct {
	Code          string
	Message       string
	OriginalError string
}

func ProduceErrorEditionPost(errorMsg string) share.BusinessError {
	return &EditionPostFailed{
		Code:          "P-102",
		Message:       "Post cannot edit",
		OriginalError: errorMsg,
	}
}

func (error *EditionPostFailed) GetCodeError() string {
	return error.Code
}

func (error *EditionPostFailed) Error() string {
	return error.Message
}

func (error *EditionPostFailed) String() string {
	return fmt.Sprintf("EditionPostFailed %s - %s [%s]", error.Code, error.Message, error.OriginalError)
}

type EliminationPostFailed struct {
	Code          string
	Message       string
	OriginalError string
}

func ProduceErrorEliminationPost(errorMsg string) share.BusinessError {
	return &EliminationPostFailed{
		Code:          "P-103",
		Message:       "Post cannot unpublish",
		OriginalError: errorMsg,
	}
}

func (error *EliminationPostFailed) GetCodeError() string {
	return error.Code
}

func (error *EliminationPostFailed) Error() string {
	return error.Message
}

func (error *EliminationPostFailed) String() string {
	return fmt.Sprintf("EliminationPostFailed %s - %s [%s]", error.Code, error.Message, error.OriginalError)
}

type PostNotFound struct {
	Code          string
	Message       string
	OriginalError string
}

func ProduceErrorPostNotFound(postId string) share.BusinessError {
	return &PostNotFound{
		Code:          "P-104",
		Message:       fmt.Sprintf("Post with Id %s not found", postId),
		OriginalError: "Post service not found Post",
	}
}

func (error *PostNotFound) GetCodeError() string {
	return error.Code
}

func (error *PostNotFound) Error() string {
	return error.Message
}

func (error *PostNotFound) String() string {
	return fmt.Sprintf("PostNotFound %s - %s [%s]", error.Code, error.Message, error.OriginalError)
}
