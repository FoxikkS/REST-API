package HttpServer

import (
	"REST-API-pet-proj/Internal/Storage"
	"REST-API-pet-proj/Internal/Storage/Sqlite"
	"REST-API-pet-proj/Pkg"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func Init() {
	cfg := Pkg.InitConfigParser()
	storage, err := Sqlite.InitStorage(cfg.Default.StoragePath)
	router := InitRouter(storage)
	Log := Pkg.SetupLogger(cfg.Default.Env)
	addr := fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port)

	Storage.Storage = *storage
	if err != nil {
		Log.Error("Error initializing sqlite storage", "error", err)
		os.Exit(1)
	}

	Log.Info("Starting server",
		slog.String("Environment", cfg.Default.Env),
		slog.String("Storage Path", cfg.Default.StoragePath),
		slog.Bool("Debug mod", cfg.Default.DebugMod),
		slog.String("status", "initializing"),
	)

	Log.Info("Listening on", slog.String("addr", addr))
	err = http.ListenAndServe(addr, router)
	if err != nil {
		Log.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
