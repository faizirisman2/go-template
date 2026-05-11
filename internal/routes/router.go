package routes

import (
	"template/go-template/internal/wire"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, wire *wire.Wire) {

	// expose docs folder
	app.Static("/docs", "./docs")
	//app.Get("/swagger/*", swagger.HandlerDefault)

	v1 := app.Group("/v1")
	v1Health := v1.Group("/health")
	v1Health.Get("/", wire.HealthHandler.Ping)

}
