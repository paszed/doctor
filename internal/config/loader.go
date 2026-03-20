package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load() Config {
	var cfg Config

	// try local file first
	data, err := os.ReadFile("doctor.yaml")
	if err != nil {
		// fallback: empty config
		return cfg
	}

	_ = yaml.Unmarshal(data, &cfg)
	return cfg
}
