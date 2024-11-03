package main

import (
	"log"

	"github.com/RTCLIF/topic-service/internal/app"
)

func main() {

	app, err := app.NewApp()

	if err != nil {
		log.Fatalf("failed to init app: %s", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}
