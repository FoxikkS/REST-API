package main

import (
	"REST-API-pet-proj/Internal/Handlers"
	"REST-API-pet-proj/Internal/Http-server/Handlers/Api/UserApi"
	"REST-API-pet-proj/Internal/Storage"
	"REST-API-pet-proj/Internal/Storage/Sqlite"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	Cfg := Handlers.InitConfigParser()
	addr := fmt.Sprintf("%s:%d", Cfg.Server.Address, Cfg.Server.Port)
	Log := Handlers.SetupLogger(Cfg.Default.Env)
	storage, err := Sqlite.InitStorage(Cfg.Default.StoragePath)
	Storage.Storage = *storage
	if err != nil {
		Log.Error("Error initializing sqlite storage", "error", err)
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	//router.Use(middleware.Logger)
	//router.Use(middleware.RealIP)

	router.Route("/api", func(r chi.Router) {
		r.Post("/register", UserApi.UserRegistrationHandler(storage))
		//r.Post("/login", UserApi.UserLoginHandler(storage))
		//r.Get("/user/{id}", UserApi.GetUserHandler(storage))
	})

	Log.Info("Starting server",
		slog.String("Environment", Cfg.Default.Env),
		slog.String("Storage Path", Cfg.Default.StoragePath),
		slog.Bool("Debug mod", Cfg.Default.DebugMod),
		slog.String("status", "initializing"),
	)

	Log.Info("Listening on", slog.String("addr", addr))
	err = http.ListenAndServe(addr, router)
	if err != nil {
		Log.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
