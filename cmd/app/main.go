package main

import (
	"os"
	"os/signal"
	"syscall"
	"template/go-template/internal/app"

	"github.com/gofiber/fiber/v2/log"
)

// @title Task API
// @version 1.0
// @description Go Template
// @host localhost:8080
// @BasePath /

func main() {

	application, err := app.BuildApplication()
	if err != nil {
		log.Error(err)
		return
	}

	go func() {
		if err = application.Run(); err != nil {
			log.Error(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	log.Info("Shutdown signal received")

	application.Shutdown()
}
