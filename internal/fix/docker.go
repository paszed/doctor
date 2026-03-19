package fix

import (
	"fmt"
	"os/exec"
)

func FixDocker() error {
	fmt.Println("[FIX] docker")

	// check if already running
	check := exec.Command("docker", "info")
	if err := check.Run(); err == nil {
		fmt.Println("✓ already running")
		return nil
	}

	fmt.Println("→ starting Docker...")

	cmd := exec.Command("open", "-a", "Docker")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start Docker\n→ try manually: open -a Docker")
	}

	fmt.Println("✓ Docker started")
	return nil
}
func init() {
	Register("docker", FixDocker)
}
