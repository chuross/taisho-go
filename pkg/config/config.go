package config

import "os"

var (
	instance *Config = &Config{
		Line: &LineConfig{
			ClientSecret:      os.Getenv("TAISHO_LINE_SECRET"),
			ClientAccessToken: os.Getenv("TAISHO_LINE_ACCESS_TOKEN"),
		},
	}
)

type Config struct {
	Line *LineConfig
}

type LineConfig struct {
	ClientSecret      string
	ClientAccessToken string
}

func Get() *Config {
	return instance
}
