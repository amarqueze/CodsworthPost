package firebasedbprovider

import (
	"outergeekhub.com/codsworthpost/domain/model/publication"
	"outergeekhub.com/codsworthpost/domain/model/share"
)

type AuthorServiceAdapter struct {
	Version string
}

func NewAuthorService() publication.AuthorService {
	return &AuthorServiceAdapter{Version: "1.0.0"}
}

func (adapter *AuthorServiceAdapter) CreateAuthor(author *publication.Author) (*publication.Author, share.BusinessError) {
	return author, nil
}

func (adapter *AuthorServiceAdapter) EditAuthor(author *publication.Author) (*publication.Author, share.BusinessError) {
	return author, nil
}

func (adapter *AuthorServiceAdapter) FindAuthor(authorId string) (*publication.Author, share.BusinessError) {
	return &publication.Author{Fullname: "Alan Marquez", Description: "Great Author"}, nil
}
