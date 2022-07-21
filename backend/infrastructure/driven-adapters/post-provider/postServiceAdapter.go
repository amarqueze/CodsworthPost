package post_Provider

import (
	"outergeekhub.com/codsworthpost/domain/model/publication"
)

type PostServiceAdapter struct {
	Version string
}

func NewPostService() publication.PostService {
	return PostServiceAdapter{Version: "1.0.0"}
}

func (adapter PostServiceAdapter) NewPublish(post *publication.Post) (*publication.Post, error) {
	return post, nil
}

func (adapter PostServiceAdapter) EditPublish(post *publication.Post) (*publication.Post, error) {
	return post, nil
}

func (adapter PostServiceAdapter) UnPublish(postId string) (*publication.Post, error) {
	postRemoved := publication.NewPost().Title("Any Great Post").Build()
	return postRemoved, nil
}
