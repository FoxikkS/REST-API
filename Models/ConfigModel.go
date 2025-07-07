package Models

import "time"

type Config struct {
	Default struct {
		Env         string `env:"ENV"`
		StoragePath string `env:"STORAGE_PATH"`
		DebugMod    bool   `env:"DEBUG_MOD"`
	}

	Server struct {
		Address     string        `env:"SERVER_ADDRESS"`
		Port        int           `env:"SERVER_PORT"`
		Timeout     time.Duration `env:"SERVER_TIMEOUT"`
		IdleTimeout time.Duration `env:"SERVER_IDLE_TIMEOUT"`
	}
}
