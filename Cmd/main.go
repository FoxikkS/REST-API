package main

import (
	"REST-API-pet-proj/Internal/Handlers"
	"log/slog"
)

func main() {
	cfg := Handlers.InitConfigParser()
	log := Handlers.SetupLogger(cfg.Default.Env)

	log.Info("Starting server",
		slog.String("Environment", cfg.Default.Env),
		slog.String("Storage Path", cfg.Default.StoragePath),
		slog.Bool("Debug mod", cfg.Default.DebugMod),
		slog.String("status", "initializing"),
	)
	log.Debug("Test debug")
}
