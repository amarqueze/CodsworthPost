package publication

type PostService interface {
	NewPublish(post *Post) (*Post, error)
	EditPublish(post *Post) (*Post, error)
	UnPublish(postId string) (*Post, error)
}

type CategoryService interface {
	CreateCategory(category *Category) (*Category, error)
}

type UploadService interface {
	UploadMedia() (string, error)
}

type AuthorService interface {
	CreateAuthor(author *Author) (*Author, error)
	EditAuthor(author *Author) (*Author, error)
}
