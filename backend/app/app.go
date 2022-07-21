package app

import (
	"fmt"

	"outergeekhub.com/codsworthpost/domain/model/publication"
	"outergeekhub.com/codsworthpost/domain/usecases"
)

type App struct {
	Title   string
	Name    string
	Version string

	Components *Components
}

type Components struct {
	PostService   *publication.PostService
	PostPublisher *usecases.PostPublisher
}

func NewApplication(title, name, version string) *App {
	app := &App{Title: title, Name: name, Version: version}
	app.Components = Setup()
	return app
}

func initComponents(
	postService publication.PostService,
	postPublisher usecases.PostPublisher) *Components {

	return &Components{PostService: &postService, PostPublisher: &postPublisher}
}

func (app *App) Run() {
	fmt.Printf("Run Application: %s v%s \n", app.Title, app.Version)
	post := publication.NewPost().
		Title("Any Great Post").
		Summary("You will feel great when you read this Post").
		Build()

	postPublisher := app.Components.PostPublisher
	postPublisher.Publish(post)
}
