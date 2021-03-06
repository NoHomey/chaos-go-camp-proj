package main

import (
	"log"

	"github.com/NoHomey/chaos-go-camp-proj/env"
	"github.com/NoHomey/chaos-go-camp-proj/expose/fxmod/app"
	"github.com/NoHomey/chaos-go-camp-proj/expose/fxmod/user"
	"go.uber.org/fx"
)

func main() {
	err := env.Load(".", "vars")
	if err != nil {
		log.Fatalln("Failed to load env vars, error: ", err.Error())
	}
	fxApp := fx.New(fx.Options(
		app.Module,
		user.Module,
	))
	fxApp.Run()
}
