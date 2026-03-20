package detect

import "os"

type Project struct {
	HasGo     bool
	HasNode   bool
	HasDocker bool
}

func Detect() Project {
	return Project{
		HasGo:     fileExists("go.mod"),
		HasNode:   fileExists("package.json"),
		HasDocker: fileExists("Dockerfile") || fileExists("docker-compose.yml"),
	}
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}
