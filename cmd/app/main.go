package main

import (
	"context"
	"log"

	"github.com/RTCLIF/topic-service/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)

	if err != nil {
		log.Fatalf("failed to init app: %s", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}
