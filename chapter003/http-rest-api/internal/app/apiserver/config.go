package apiserver

import "src/myFirstGoProject/chapter003/http-rest-api/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
	//DatabaseURL string `toml:"database_url"`
	//SessionKey  string `toml:"session_key"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8034",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
