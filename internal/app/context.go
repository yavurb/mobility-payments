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

	authApplication "github.com/yavurb/mobility-payments/internal/auth/application"
	authAdapters "github.com/yavurb/mobility-payments/internal/auth/infrastructure/adapters"
	authHTTPUI "github.com/yavurb/mobility-payments/internal/auth/infrastructure/ui/http"

	userApplication "github.com/yavurb/mobility-payments/internal/users/application"
	userRepositoryAdapter "github.com/yavurb/mobility-payments/internal/users/infrastructure/adapters/repository"

	paymentApplication "github.com/yavurb/mobility-payments/internal/payments/application"
	paymentRepositoryAdapter "github.com/yavurb/mobility-payments/internal/payments/infrastructure/adapters/repository"
	paymentHTTPUI "github.com/yavurb/mobility-payments/internal/payments/infrastructure/ui/http"
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

	userRepository := userRepositoryAdapter.NewUserRepository(app.connpool)
	userUsecase := userApplication.NewUserUsecase(userRepository)
	authUsecase := authApplication.NewAuthUsecase(
		userUsecase,
		authAdapters.NewAuthPasswordHasher(),
		authAdapters.NewAuthTokenManager(app.config.HttpAuth.JWTSecret),
	)
	authHTTPUI.NewAuthRouter(e, authUsecase)

	paymentRepository := paymentRepositoryAdapter.NewPaymentsRepository(app.connpool)
	paymentUsecase := paymentApplication.NewPaymentsUsecase(paymentRepository, userUsecase)
	paymentHTTPUI.NewPaymentsRouter(e, paymentUsecase)

	return e
}
