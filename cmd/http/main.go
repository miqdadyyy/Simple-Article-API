package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/google/logger"

	"github.com/miqdadyyy/jetdevs-assesment/pkg/config"
)

func main() {
	fatalChan := make(chan error)
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "local"
	}
	cfg := config.GetConfig(environment)

	app := fiber.New()

	// Initialize all routes
	InitializeRoutes(app)

	// Listen for application
	err := app.Listen(fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port))
	if err != nil {
		fatalChan <- err
	}

	// Graceful shutdown
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		logger.Error("SigTerm Detected")
	case err := <-fatalChan:
		logger.Error("Failed to run application", err.Error())
	}
}
