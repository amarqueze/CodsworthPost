package web

import (
	"outergeekhub.com/codsworthpost/domain/model/publication"
	"outergeekhub.com/codsworthpost/domain/usecases"
)

type HomeController struct {
	Publisher usecases.PostPublisher
}

func NewHomeController(publisher usecases.PostPublisher) HomeController {
	return HomeController{Publisher: publisher}
}

func (controller HomeController) ReceiveRequest() {
	post := publication.NewPost("en-EN").
		Title("Any Great Post").
		Summary("You will feel great when you read this Post").
		Build()
	controller.Publisher.Publish(post)
}

func health() {

}
