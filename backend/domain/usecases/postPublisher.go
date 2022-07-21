package usecases

import (
	"fmt"

	"outergeekhub.com/codsworthpost/domain/model/publication"
)

type PostPublisher struct {
	postService publication.PostService
}

func NewPostPublisher(postService publication.PostService) PostPublisher {
	return PostPublisher{postService}
}

func (pub *PostPublisher) Publish(post *publication.Post) *publication.Post {
	post, error := pub.postService.NewPublish(post)
	if error != nil {
		fmt.Println("Publish has failed!")
		return post
	}

	fmt.Println("Publish success!")
	return post
}
