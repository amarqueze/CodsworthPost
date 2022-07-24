package app

import (
	"fmt"

	"outergeekhub.com/codsworthpost/domain/model/publication"
	"outergeekhub.com/codsworthpost/domain/usecases"
	"outergeekhub.com/codsworthpost/infrastructure/entry-points/web"
)

type App struct {
	Title   string
	Name    string
	Version string

	Components *Components
}

type Components struct {
	PostService    *publication.PostService
	PostPublisher  *usecases.PostPublisher
	HomeController *web.HomeController
}

func NewApplication(title, name, version string) *App {
	app := &App{Title: title, Name: name, Version: version}
	app.Components = Setup()
	return app
}

func initComponents(
	postService publication.PostService,
	postPublisher usecases.PostPublisher,
	homeController web.HomeController) *Components {

	return &Components{
		PostService:    &postService,
		PostPublisher:  &postPublisher,
		HomeController: &homeController,
	}
}

func (app *App) Run() {
	fmt.Printf("Run Application: %s v%s \n", app.Title, app.Version)
	controller := app.Components.HomeController
	controller.ReceiveRequest()
}
