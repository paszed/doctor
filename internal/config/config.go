package config

var Current Config

type Config struct {
	Ignore []string `yaml:"ignore"`
	Ports  []int    `yaml:"ports"`
}

func merge(global, local Config) Config {
	out := global

	// override ignore if local has it
	if len(local.Ignore) > 0 {
		out.Ignore = local.Ignore
	}

	// override ports if local has it
	if len(local.Ports) > 0 {
		out.Ports = local.Ports
	}

	return out
}
