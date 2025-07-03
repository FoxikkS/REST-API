package structure

import "time"

type Config struct {
	Default struct {
		Env         string `yaml:"Env"`
		StoragePath string `yaml:"StoragePath"`
		DebugMod    bool   `yaml:"DebugMod"`
	} `yaml:"DEFAULT"`

	Server struct {
		Address     string        `yaml:"Address"`
		Port        int           `yaml:"Port"`
		Timeout     time.Duration `yaml:"Timeout"`
		IdleTimeout time.Duration `yaml:"IdleTimeout"`
	} `yaml:"SERVER"`
}
