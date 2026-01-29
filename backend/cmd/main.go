package main

import (
	"log"
	"runtime/debug"

	"github.com/spyatachkov/green-api/backend/internal/app"
	"github.com/spyatachkov/green-api/backend/internal/config"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("application panic: %v\n%s", err, debug.Stack())
			log.Fatal("Application crashed")
		}
	}()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	application := app.New(cfg)

	if err := application.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
