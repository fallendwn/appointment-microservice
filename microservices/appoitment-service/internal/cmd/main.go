package main

import "github.com/fallendwn/appointment/appointment-service/internal/app"

func main() {
	cfg := app.NewConfig()
	app.Run(cfg)
}
