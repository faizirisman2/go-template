package app

import (
	"template/go-template/internal/config"
	"template/go-template/internal/dependency"
	"template/go-template/internal/handler"
	"template/go-template/internal/routes"
	"template/go-template/internal/wire"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type App struct {
	Cfg        *config.Config
	fiber      *fiber.App
	Wire       *wire.Wire
	Dependency *dependency.Dependency
}

func BuildApplication() (*App, error) {

	fiberApp := fiber.New()

	healthHandler := handler.NewHealthHandler()

	wire := &wire.Wire{
		HealthHandler: healthHandler,
	}

	routes.SetupRouter(fiberApp, wire)

	dep := dependency.InitializeDependency()

	return &App{
		Cfg:        dep.Cfg,
		fiber:      fiberApp,
		Wire:       wire,
		Dependency: dep,
	}, nil
}

func (app *App) Run() error {

	// =====================================================
	// Start Server
	// =====================================================

	address := ":" + app.Cfg.App.Port

	log.Debugf(
		"starting %s on %s",
		app.Cfg.App.Name,
		address,
	)

	if err := app.fiber.Listen(address); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (app *App) Shutdown() {

	log.Debugf("shutting down app %s", app.Cfg.App.Name)
	if err := app.fiber.Shutdown(); err != nil {
		log.Error(err)
	}

	log.Info("closing postgres")
	app.Dependency.Postgres.Close()
	log.Info("shutdown completed")
}
