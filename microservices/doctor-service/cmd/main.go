package main

import "github.com/fallendwn/appointment/doctor-service/internal/app"

func main() {
	cfg := app.NewConfig()
	app.Run(cfg)
}
