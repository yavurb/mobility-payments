package app

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	appconfig "github.com/yavurb/mobility-payments/config/app_config"
	"github.com/yavurb/mobility-payments/config/app_config/loglevel"
	"github.com/yavurb/mobility-payments/internal/app/mods"
)

type App struct {
	config   *appconfig.Config
	connpool *pgxpool.Pool
}

func NewApp(config *appconfig.Config) *App {
	ctx := context.Background()

	connpool, err := pgxpool.New(ctx, config.Database.URI)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	return &App{
		config:   config,
		connpool: connpool,
	}
}

func (app *App) NewHttpContext() *echo.Echo {
	e := echo.New()

	var logLevel log.Lvl

	switch app.config.LogLevel {
	case loglevel.Debug:
		logLevel = log.DEBUG
	case loglevel.Info:
		logLevel = log.INFO
	case loglevel.Warn:
		logLevel = log.WARN
	case loglevel.Error:
		logLevel = log.ERROR
	}

	e.Logger.SetLevel(logLevel)
	e.Validator = mods.NewAppValidator()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	e.GET("/health", func(c echo.Context) error { return c.String(200, "Healthy!") })

	return e
}
