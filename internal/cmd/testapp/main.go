package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	cfgP "github.com/Cwby333/testapp/internal/config"
	"github.com/Cwby333/testapp/internal/infra/presentation/http"
	"github.com/Cwby333/testapp/internal/infra/repository/postgresql"
)

func main() {
	logger := newLogger()
	_ = logger

	cfg := cfgP.Load()
	fmt.Println(cfg)

	cfgPg := postgresql.Config{
		Host: cfg.Postgres.Host,
		Port: cfg.Postgres.Port,
		User: cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		Database: cfg.Postgres.DB,
		MaxConns: cfg.Postgres.MaxConns,
		MinConns: cfg.Postgres.MinConns,
		ConnIdle: cfg.Postgres.MaxIdleConn,
		ConnLife: cfg.Postgres.MaxLifetimeConn,
	}
	pg, err := postgresql.New(context.Background(), cfgPg)
	if err != nil {
		panic(err)
	}

	err = pg.Insert(context.Background(), 1)
	v, err := pg.Get(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

	handler := http.New(pg)
	server := http.NewServer(*handler)

	fmt.Println(server)

	log.Fatal(server.ListenAndServe())
}

func newLogger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	logger := slog.New(handler)

	return logger
} 