package app

import (
	"github.com/google/wire"
	"outergeekhub.com/codsworthpost/domain/usecases"
	post_Provider "outergeekhub.com/codsworthpost/infrastructure/driven-adapters/post-provider"
	"outergeekhub.com/codsworthpost/infrastructure/entry-points/web"
)

var PostProvider = wire.NewSet(post_Provider.NewPostService)
var PostPublisher = wire.NewSet(usecases.NewPostPublisher)
var HomeController = wire.NewSet(web.NewHomeController)
var APIComponent = wire.NewSet(initComponents)
