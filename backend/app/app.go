package app

import (
	"fmt"

	apirest "outergeekhub.com/codsworthpost/infrastructure/entry-points/api_rest"
)

type App struct {
	Title   string
	Name    string
	Version string

	Components *Components
}

type Components struct {
	PublishPostPostController *apirest.PublishPostPostController
}

func NewApplication(title, name, version string) *App {
	app := &App{Title: title, Name: name, Version: version}
	app.Components = Setup()
	return app
}

func initComponents(
	publishPostPostController apirest.PublishPostPostController,
) *Components {

	return &Components{
		PublishPostPostController: &publishPostPostController,
	}
}

func (app *App) Run() {
	fmt.Printf("Run Application: %s v%s \n", app.Title, app.Version)
	controller := app.Components.PublishPostPostController
	controller.ReceiveRequest()
}
