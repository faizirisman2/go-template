package dependency

import (
	"template/go-template/internal/config"
	"template/go-template/internal/database"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependency struct {
	Cfg      *config.Config
	Postgres *pgxpool.Pool
}

func InitializeDependency() *Dependency {
	cfg := config.LoadConfig()

	dbCon, err := database.NewPostgresConnection(&cfg.DB)
	if err != nil {
		log.Error(err)
	}

	return &Dependency{
		Cfg: &config.Config{
			App: cfg.App,
			DB:  cfg.DB,
		},
		Postgres: dbCon,
	}
}
