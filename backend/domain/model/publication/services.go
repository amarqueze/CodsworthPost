package publication

import "outergeekhub.com/codsworthpost/domain/model/share"

type PostService interface {
	NewPublish(post *Post) (*Post, share.BusinessError)
	EditPublish(post *Post) (*Post, share.BusinessError)
	UnPublish(postId string) (*Post, share.BusinessError)

	FindPost(postId string) (*Post, share.BusinessError)
	FindPostsByCriteria() (*[]Post, share.BusinessError)
}

type CategoryService interface {
	CreateCategory(category *Category) (*Category, share.BusinessError)
	FindCategory(categoryName string) (*Category, share.BusinessError)
}

type UploadService interface {
	UploadMedia() (string, share.BusinessError)
}

type AuthorService interface {
	CreateAuthor(author *Author) (*Author, share.BusinessError)
	EditAuthor(author *Author) (*Author, share.BusinessError)
	FindAuthor(authorId string) (*Author, share.BusinessError)
}
