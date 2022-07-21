package main

import "outergeekhub.com/codsworthpost/app"

func main() {
	app := app.NewApplication("Outer Geek Hub", "codsworthpost", "1.0.0")
	app.Run()
}
