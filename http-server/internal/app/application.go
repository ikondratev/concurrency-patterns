package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	netHttp "net/http"

	"github.com/kilia/http-example/internal/http"
	"github.com/kilia/http-example/internal/settings"
)

type Server interface {
	Start(chan<- error)
	Stop(context.Context)error
}

type Application struct {
	settings *settings.Settings
	server Server
}

func New(env string) (*Application, error) {
	// Settings
	settings, err := settings.Load(env)
	if err != nil {
		return nil, fmt.Errorf("Settings error: %w", err)
	}

	//Server
	server := http.NewServer(settings)

	return &Application{
		settings: settings,
		server:   server,
	}, nil
}

func (a *Application) Run() error {
	chErr := make(chan error, 1)
	go a.server.Start(chErr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-chErr:
		if !errors.Is(err, netHttp.ErrServerClosed) {
			return err
		}
	case sig := <- quit:
		fmt.Println("Server received signal:", sig)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 
		time.Duration(a.settings.Server.ShutdownTimeoutSecond) * time.Second)
	defer cancel()

	if err := a.server.Stop(ctx); err != nil {
		return err
	}

	return nil
}
