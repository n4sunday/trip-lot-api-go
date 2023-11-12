package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"trip-lot-api/config"
	"trip-lot-api/internal/adapters/handler"
	"trip-lot-api/routes"
	"trip-lot-api/store/postgres"
	"trip-lot-api/store/postgres/migrations"

	userRepo "trip-lot-api/internal/adapters/repository/user/postgres"
	userService "trip-lot-api/internal/core/user"

	"github.com/labstack/echo/v4"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	e := echo.New()
	e.GET("/api/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server ready")
	})

	// Environment config
	appConfig := config.AppConfig()

	// Init Database
	postgres.InitDB(appConfig.DatabaseUrl)
	db := postgres.GetDB()
	defer postgres.CloseDB()

	// Init Migrations
	migrations.Migrate(db)

	// Init User
	userRepo := userRepo.NewUserPgRepository(db)
	userService := userService.NewUserService(userRepo)
	userHandler := handler.NewUserHTTPHandler(userService)

	// Routes
	routes.NewUserRoute(e, userHandler)

	// Start server
	go func() {
		if err := e.Start(":" + appConfig.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
