package Pkg

import (
	"REST-API-pet-proj/Models"
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

var ConfigPath = os.Getenv("CONFIG_PATH")

func checkConfig() bool {
	if ConfigPath == "" {
		slog.Error("CONFIG_PATH environment variable not set")
		return false
	}
	slog.Info("CONFIG_PATH: " + ConfigPath)

	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		slog.Error("CONFIG_PATH does not exist")
		return false
	}
	return true
}

func InitConfigParser() Models.Config {
	var Cfg Models.Config
	if checkConfig() == true {
		slog.Info("Reading config")
		err := cleanenv.ReadConfig(ConfigPath, &Cfg)
		if err != nil {
			slog.Error("Error reading config file, %s", err)
		}
		slog.Info("Config is parsed!")
	}
	return Cfg
}
