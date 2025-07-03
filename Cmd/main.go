package main

import (
	"REST-API-pet-proj/Internal/Handlers"
	"REST-API-pet-proj/Internal/Storage/Sqlite"
	"log/slog"
	"os"
)

func main() {
	cfg := Handlers.InitConfigParser()
	log := Handlers.SetupLogger(cfg.Default.Env)
	storage, err := Sqlite.InitStorage(cfg.Default.StoragePath)
	if err != nil {
		log.Error("Error initializing sqlite storage", "error", err)
		os.Exit(1)
	}

	_ = storage

	log.Info("Starting server",
		slog.String("Environment", cfg.Default.Env),
		slog.String("Storage Path", cfg.Default.StoragePath),
		slog.Bool("Debug mod", cfg.Default.DebugMod),
		slog.String("status", "initializing"),
	)
}
