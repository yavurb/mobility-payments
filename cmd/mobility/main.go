package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	appconfig "github.com/yavurb/mobility-payments/config/app_config"
	"github.com/yavurb/mobility-payments/internal/app"
)

func main() {
	cfg := appconfig.LoadConfig()
	app := app.NewApp(cfg)
	httpServer := app.NewHttpContext()
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	go func() {
		if err := httpServer.Start(addr); err != nil && err != http.ErrServerClosed {
			httpServer.Logger.Fatal("shutting down the server", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		httpServer.Logger.Fatal(err)
	}
}
