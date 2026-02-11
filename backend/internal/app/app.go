package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tdzdslippen/suplan/internal/storage/db"
	httptransport "github.com/tdzdslippen/suplan/internal/transport/http"
)

type App struct {
	cfg Config
	db  *db.DB
}

func New(ctx context.Context) (*App, error) {
	cfg, err := Load()
	if err != nil {
		return nil, err
	}
	database, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return &App{cfg: cfg, db: database}, nil

}

func (a *App) Run() error {
	srv := &http.Server{
		Addr:              a.cfg.HTTPAddr,
		Handler:           httptransport.NewRouter(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	errCh := make(chan error, 1)
	go func() {
		log.Printf("api listening on %s (env=%s)", a.cfg.HTTPAddr, a.cfg.Env)
		errCh <- srv.ListenAndServe()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-quit:
		log.Printf("shutdown signal: %s", sig.String())
	case err := <-errCh:
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = srv.Shutdown(ctx)
	return nil
}
