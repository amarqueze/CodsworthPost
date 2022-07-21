package apirest

import (
	"outergeekhub.com/codsworthpost/domain/model/publication"
	"outergeekhub.com/codsworthpost/domain/usecases"
)

type HomeController struct {
	Publisher usecases.PostPublisher
}

func (controller HomeController) ReceiveRequest() {
	post := publication.NewPost().Title("Fisrt Post").Build()
	controller.Publisher.Publish(post)
}
