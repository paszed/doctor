package config

var Current Config

type Config struct {
	Ignore []string `yaml:"ignore"`
	Ports  []int    `yaml:"ports"`
}
