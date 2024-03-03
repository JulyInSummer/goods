package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/nats-io/nats.go"
	"goods_project/internal/api"
	"goods_project/internal/cache"
	"goods_project/internal/config"
	"goods_project/internal/service"
	"goods_project/internal/storage/pg-storage"
)

const (
	local = "local"
	dev   = "dev"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	defer nc.Close()

	ctx := context.Background()
	cfg := config.NewConfig()
	logger := LoadLogger(cfg.App.Env)
	db := pg_storage.NewPGStorage(cfg, logger)
	redis := cache.NewCache(cfg)
	repo := service.NewService(logger, db, redis, nc)
	serv := api.NewServer(cfg, repo, logger, ctx)

	serv.Start()
}

func LoadLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case local:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case dev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
