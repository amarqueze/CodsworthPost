package app

import (
	"github.com/google/wire"
	"outergeekhub.com/codsworthpost/domain/usecases"
	firebasedbprovider "outergeekhub.com/codsworthpost/infrastructure/driven-adapters/firebaseDb-provider"
	apirest "outergeekhub.com/codsworthpost/infrastructure/entry-points/api_rest"
)

//** Driven Adapters
var PostProvider = wire.NewSet(firebasedbprovider.NewPostService)

//** Uses Cases
var PostPublisher = wire.NewSet(usecases.NewPostPublisher)

//** Entry Points
var PublishPostPostController = wire.NewSet(apirest.NewPublishPostPostController)

var APIComponent = wire.NewSet(initComponents)
