package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Load() Config {
	global := loadFile(globalPath())
	local := loadFile("doctor.yaml")

	return merge(global, local)
}

func loadFile(path string) Config {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg
	}

	_ = yaml.Unmarshal(data, &cfg)
	return cfg
}

func globalPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".doctor.yaml")
}
