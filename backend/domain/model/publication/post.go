package publication

import (
	"time"
)

type Post struct {
	id           string
	title        string
	summary      string
	content      string
	primaryImage string
	category     Category
	tags         []string
	dateCreated  string
	dateEdition  string
	IdWriter     string
}

func NewPost() *NewPostBuilder {
	newPostBuilder := new(NewPostBuilder)
	newPostBuilder.post.dateCreated = time.Now().String()
	return newPostBuilder
}

type NewPostBuilder struct {
	post Post
}

func (builder *NewPostBuilder) Title(title string) *NewPostBuilder {
	builder.post.title = title
	return builder
}

func (builder *NewPostBuilder) Summary(summary string) *NewPostBuilder {
	builder.post.summary = summary
	return builder
}

func (builder *NewPostBuilder) Content(content string) *NewPostBuilder {
	builder.post.content = content
	return builder
}

func (builder *NewPostBuilder) PrimaryImage(primaryImage string) *NewPostBuilder {
	builder.post.primaryImage = primaryImage
	return builder
}

func (builder *NewPostBuilder) Category(category Category) *NewPostBuilder {
	builder.post.category = category
	return builder
}

func (builder *NewPostBuilder) AddTags(tags []string) *NewPostBuilder {
	builder.post.tags = tags
	return builder
}

func (builder *NewPostBuilder) Writer(idWriter string) *NewPostBuilder {
	builder.post.IdWriter = idWriter
	return builder
}

func (builder *NewPostBuilder) Build() *Post {
	return &builder.post
}

func EditPost(id string) *EditPostBuilder {
	editPostBuilder := new(EditPostBuilder)
	editPostBuilder.post.id = id
	editPostBuilder.post.dateEdition = time.Now().String()
	return editPostBuilder
}

type EditPostBuilder struct {
	post Post
}

func (builder *EditPostBuilder) Content(content string) *EditPostBuilder {
	builder.post.content = content
	return builder
}

func (builder *EditPostBuilder) AddTags(tags []string) *EditPostBuilder {
	builder.post.tags = tags
	return builder
}

func (builder *EditPostBuilder) Build() *Post {
	return &builder.post
}
