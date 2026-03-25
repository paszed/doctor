package fix

import (
	"fmt"
	"os/exec"
)

func FixDockerCompose(args []string) error {
	fmt.Println("[FIX] docker-compose")
	cmd := exec.Command("docker-compose", "--version")
	if err := cmd.Run(); err == nil {
		fmt.Println("✓ Docker Compose is already installed")
		return nil
	}
	fmt.Println("→ docker-compose not found. Please install manually:")
	fmt.Println("   https://docs.docker.com/compose/install/")
	return nil
}

func init() {
	Register("docker-compose", FixDockerCompose)
}
