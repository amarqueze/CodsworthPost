package post_Provider

import (
	"outergeekhub.com/codsworthpost/domain/model/publication"
	"outergeekhub.com/codsworthpost/domain/model/share"
)

type PostServiceAdapter struct {
	Version string
}

func NewPostService() publication.PostService {
	return &PostServiceAdapter{Version: "1.0.0"}
}

func (adapter PostServiceAdapter) NewPublish(post *publication.Post) (*publication.Post, share.BusinessError) {
	return post, nil
}

func (adapter PostServiceAdapter) EditPublish(post *publication.Post) (*publication.Post, share.BusinessError) {
	return post, nil
}

func (adapter PostServiceAdapter) UnPublish(postId string) (*publication.Post, share.BusinessError) {
	postRemoved := publication.Post{Id: postId, Title: "Any Great Post", Lang: "en-En"}
	return &postRemoved, nil
}

func (adapter PostServiceAdapter) FindPost(postId string) (*publication.Post, share.BusinessError) {
	return nil, nil
}

func (adapter PostServiceAdapter) FindPostsByCriteria() (*[]publication.Post, share.BusinessError) {
	return nil, nil
}
