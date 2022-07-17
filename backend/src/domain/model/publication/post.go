package publication

import "time"

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

func NewPost() newPostBuilder {
	newPostBuilder := new(newPostBuilder)
	newPostBuilder.post.dateCreated = time.Now().String()
	return *newPostBuilder
}

type newPostBuilder struct {
	post Post
}

func (builder newPostBuilder) Title(title string) newPostBuilder {
	builder.post.title = title
	return builder
}

func (builder newPostBuilder) Summary(summary string) newPostBuilder {
	builder.post.summary = summary
	return builder
}

func (builder newPostBuilder) Content(content string) newPostBuilder {
	builder.post.content = content
	return builder
}

func (builder newPostBuilder) PrimaryImage(primaryImage string) newPostBuilder {
	builder.post.primaryImage = primaryImage
	return builder
}

func (builder newPostBuilder) Category(category Category) newPostBuilder {
	builder.post.category = category
	return builder
}

func (builder newPostBuilder) AddTags(tags []string) newPostBuilder {
	builder.post.tags = tags
	return builder
}

func (builder newPostBuilder) Writer(idWriter string) newPostBuilder {
	builder.post.IdWriter = idWriter
	return builder
}

func (builder newPostBuilder) Build() Post {
	return builder.post
}

func EditPost(id string) EditPostBuilder {
	editPostBuilder := new(EditPostBuilder)
	editPostBuilder.post.id = id
	editPostBuilder.post.dateEdition = time.Now().String()
	return *editPostBuilder
}

type EditPostBuilder struct {
	post Post
}

func (builder EditPostBuilder) Content(content string) EditPostBuilder {
	builder.post.content = content
	return builder
}

func (builder EditPostBuilder) AddTags(tags []string) EditPostBuilder {
	builder.post.tags = tags
	return builder
}

func (builder EditPostBuilder) Build() Post {
	return builder.post
}
