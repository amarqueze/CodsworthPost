package apirest

import (
	"outergeekhub.com/codsworthpost/domain/model/publication"
	"outergeekhub.com/codsworthpost/domain/usecases"
)

type PublishPostPostController struct {
	Publisher usecases.PostPublisher
}

func NewPublishPostPostController(publisher usecases.PostPublisher) PublishPostPostController {
	return PublishPostPostController{Publisher: publisher}
}

func (controller *PublishPostPostController) ReceiveRequest() {
	post := publication.NewPost("en-EN").
		Title("Any Great Post").
		Summary("You will feel great when you read this Post").
		Content("This a great Content").
		Category(publication.NewCategory("Programming")).
		Writer(publication.NewAuthor("Alan Marquez E.", "Super-Duper-Hyper-Developer-Friend")).
		Build()
	controller.Publisher.Publish(post)
}
